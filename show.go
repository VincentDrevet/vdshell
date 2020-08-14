package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw"

	"github.com/abiosoft/ishell"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// AjoutShowCommande ajoutant les commandes show
func AjoutShowCommande() *ishell.Cmd {
	showCmd := &ishell.Cmd{
		Name:     "show",
		Help:     "Affiche des informations sur un elément spécifique",
		LongHelp: `Syntaxe : show <element>`,
	}
	showCmd.AddCmd(&ishell.Cmd{
		Name: "version",
		Help: "Affiche la version du shell",
		Func: func(c *ishell.Context) {

			c.Printf("version : 0.0.1\n")
		},
	})
	showCmd.AddCmd(&ishell.Cmd{
		Name: "uptime",
		Help: "Affiche la durée de fonctionnement ininterrompu du système",
		Func: func(c *ishell.Context) {
			uptime, err := host.Uptime()
			if err != nil {
				log.Println(err)
				return
			}
			var uptimeh float64 = float64(uptime) / float64(3600)
			c.Printf("Uptime : %f h\n", uptimeh)
		},
	})
	showCmd.AddCmd(&ishell.Cmd{
		Name: "hostinfo",
		Help: "Affiche les informations de la platforme",
		Func: func(c *ishell.Context) {
			titre := []string{"Hostname", "Nombre de processus", "OS", "Plateforme", "Famille de plateforme", "Version de la plateforme", "Version du noyau", "Architecture", "UUID"}
			Hostinfo, err := host.Info()
			if err != nil {
				log.Println(err)
			}
			donnees := [][]string{[]string{Hostinfo.Hostname, strconv.FormatUint(Hostinfo.Procs, 10), Hostinfo.OS, Hostinfo.Platform, Hostinfo.PlatformFamily, Hostinfo.PlatformVersion, Hostinfo.KernelVersion, Hostinfo.KernelArch, Hostinfo.HostID}}

			dessinertableau(titre, donnees)
		},
	})
	showCmd.AddCmd(&ishell.Cmd{
		Name: "volumes",
		Help: "Affiche les informations à propos des volumes",
		Func: func(c *ishell.Context) {
			titres := []string{"Chemin", "Type", "Options", "Disques", "Total", "Utilisé", "Pourcentage utilisé"}
			var donnees [][]string
			partitions, err := disk.Partitions(true)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			for _, partition := range partitions {
				usage, err := disk.Usage(partition.Mountpoint)
				if err != nil {
					log.Printf("%v", err)
				}
				donnees = append(donnees, []string{usage.Path, usage.Fstype, partition.Opts, partition.Device, strconv.FormatUint(usage.Total, 10), strconv.FormatUint(usage.Used, 10), fmt.Sprintf("%f", usage.UsedPercent)})
			}
			dessinertableau(titres, donnees)
		},
	})
	showCmd.AddCmd(&ishell.Cmd{
		Name: "disks",
		Help: "Affiche le nom des disks disponible",
		Func: func(c *ishell.Context) {
			titre := []string{"Nom", "Vendeur", "Modèle", "Numéro de série", "Contrôleur de disque", "Partition(s)"}
			var donnees [][]string
			block, err := ghw.Block()
			if err != nil {
				log.Fatal(err)
			}
			for _, disk := range block.Disks {
				var partitions []string
				for _, partition := range disk.Partitions {
					partitions = append(partitions, partition.Name+" ["+partition.Type+"]")
				}
				donnees = append(donnees, []string{disk.Name, disk.Vendor, disk.Model, disk.SerialNumber, disk.StorageController.String(), strings.Join(partitions, " ")})
			}
			dessinertableau(titre, donnees)
		},
	})
	showCmd.AddCmd(&ishell.Cmd{
		Name: "memory",
		Help: "Affiche des informations concernant la mémoire",
		Func: func(c *ishell.Context) {
			swapinfo, err := mem.SwapMemory()
			if err != nil {
				log.Printf("%v", err)
				return
			}
			meminfo, err := mem.VirtualMemory()
			if err != nil {
				log.Printf("%v", err)
				return
			}

			titres := []string{"Swap utilisé", "Swap total", "Pourcentage swap utilisé"}
			donnees := [][]string{[]string{strconv.FormatUint(swapinfo.Used, 10), strconv.FormatUint(swapinfo.Total, 10), fmt.Sprintf("%f", swapinfo.UsedPercent)}}

			dessinertableau(titres, donnees)

			titres = []string{"Mémoire disponible", "Mémoire total", "Pourcentage mémoire utilisée"}
			donnees = [][]string{[]string{strconv.FormatUint(meminfo.Used, 10), strconv.FormatUint(meminfo.Total, 10), fmt.Sprintf("%f", meminfo.UsedPercent)}}

			c.Printf("\n\n")
			dessinertableau(titres, donnees)

		},
	})
	interfacecmd := &ishell.Cmd{
		Name:     "interfaces",
		Help:     "Affiche les informations relatives aux interfaces",
		LongHelp: `Syntaxe : show interface <options>`,
	}
	interfacecmd.AddCmd(&ishell.Cmd{
		Name: "summary",
		Help: "Affiche des informations essentielles sur les interfaces",
		Func: func(c *ishell.Context) {

			titres := []string{"Index", "MTU", "Nom", "Flags", "Adresses"}
			var donnees [][]string
			var adresses []string
			interfaces, err := net.Interfaces()
			if err != nil {
				log.Printf("%v", err)
				return
			}
			for _, netinterface := range interfaces {

				for _, adresse := range netinterface.Addrs {
					adresses = append(adresses, adresse.Addr)
				}
				donnees = append(donnees, []string{strconv.Itoa(netinterface.Index), strconv.Itoa(netinterface.MTU), netinterface.Name, strings.Join(netinterface.Flags, " "), strings.Join(adresses, " ")})
			}
			dessinertableau(titres, donnees)
		},
	})
	interfacecmd.AddCmd(&ishell.Cmd{
		Name: "sockets",
		Help: "Affiche les connexions en cours et en écoute",
		Func: func(c *ishell.Context) {
			connexions, err := net.Connections("all")
			titres := []string{"Adresse Local", "Port Local", "Adresse Distante", "Port Distant", "Status"}
			var donnees [][]string
			if err != nil {
				log.Printf("%v", err)
				return
			}
			for _, connexion := range connexions {
				donnees = append(donnees, []string{connexion.Laddr.IP, strconv.FormatUint(uint64(connexion.Laddr.Port), 10), connexion.Raddr.IP, strconv.FormatUint(uint64(connexion.Raddr.Port), 10), connexion.Status})
			}
			dessinertableau(titres, donnees)
		},
	})
	showCmd.AddCmd(interfacecmd)
	return showCmd
}
