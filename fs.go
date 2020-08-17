package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"syscall"

	"github.com/jaypipes/ghw"

	"github.com/abiosoft/ishell"
)

// AjoutFSCommande Fonction ajoutant la gestion de FS
func AjoutFSCommande() *ishell.Cmd {
	FsCmd := &ishell.Cmd{
		Name:     "fs",
		Help:     "Permet d'effectuer des actions sur les systèmes de fichier",
		LongHelp: "Syntaxe : fs <option>",
	}

	FsCmd.AddCmd(&ishell.Cmd{
		Name: "mount",
		Help: "Active le périphérique distant sur la plateforme",
		Func: func(c *ishell.Context) {
			block, err := ghw.Block()
			if err != nil {
				log.Fatalf("%v", err)
			}
			// On parse l'entré de l'utilisateur
			re := regexp.MustCompile("^[a-z]+")
			match := re.FindStringSubmatch(c.Args[0])

			var retour bool = false

			for _, disk := range block.Disks {
				if match[0] == disk.Name {
					// On itère sur les partitions du disque
					for _, partition := range disk.Partitions {
						if partition.Name == c.Args[0] {
							// On vérifie le système de fichier
							if partition.Type == "vfat" {
								err := syscall.Mount("/dev/"+c.Args[0], "/mnt", "vfat", syscall.MS_NOEXEC, "")
								if err != nil {
									log.Fatalf("%v\n", err)
								}
								retour = true
							}
						}
					}
				}
			}
			if retour == false {
				log.Fatalln("Erreur lors du chargement du périphérique, le périphérique est il formaté en fat32 ?, la partition existe t-elle ?")
			}
		},
	})
	FsCmd.AddCmd(&ishell.Cmd{
		Name: "unmount",
		Help: "Désactive le périphérique distant sur la plateforme",
		Func: func(c *ishell.Context) {
			err := syscall.Unmount("/mnt", syscall.MNT_FORCE)
			if err != nil {
				log.Fatalf("%v", err)
			}
		},
	})

	FsCmd.AddCmd(&ishell.Cmd{
		Name: "ls",
		Help: "Liste le contenu d'un périphérique",
		Func: func(c *ishell.Context) {
			titre := []string{"Nom", "Dernière modification", "Permission", "Répertoire", "Taille"}
			var donnees [][]string
			err := filepath.Walk("/mnt", func(chemin string, info os.FileInfo, err error) error {
				donnees = append(donnees, []string{info.Name(), info.ModTime().String(), info.Mode().Perm().String(), strconv.FormatBool(info.IsDir()), strconv.FormatInt(info.Size(), 10)})
				return nil
			})
			if err != nil {
				log.Fatalf("%v", err)
			}
			dessinertableau(titre, donnees)

		},
	})

	return FsCmd
}
