package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/abiosoft/ishell"
)

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
			err := syscall.Mount("/dev/"+c.Args[0], "/mnt", "vfat", syscall.MS_NOEXEC, "")
			if err != nil {
				log.Fatalf("%v\n", err)
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
