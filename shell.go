package main

import (
	"fmt"

	"github.com/abiosoft/ishell"
	"github.com/shirou/gopsutil/host"
)

// TypeShell enumeration
type TypeShell int

const (
	// Restreint enumeration du type de shell
	Restreint TypeShell = iota
	// Administrateur enumeration du type de shell
	Administrateur
)

// Shell Structure de donnée permettant la construction d'un shell
type Shell struct {
	typeshell    TypeShell
	message      string
	commandes    []*ishell.Cmd
	precommandes []*ishell.Cmd
}

// NouveauShell Fonction permettant la creation de nouveau shell
func NouveauShell(shell Shell) *ishell.Shell {
	Hostinfo, err := host.Info()
	if err != nil {
		fmt.Printf("%v", err)
	}
	var nvshell *ishell.Shell
	switch shell.typeshell {
	case Restreint:
		sh := ishell.New()
		sh.Println(shell.message)
		sh.SetPrompt("[" + Hostinfo.Hostname + "] - $> ")
		for _, commande := range shell.commandes {
			sh.AddCmd(commande)
		}
		nvshell = sh

	case Administrateur:
		shelladmin := ishell.New()
		shelladmin.Println(shell.message)
		shelladmin.SetPrompt("[" + Hostinfo.Hostname + "] - #> ")
		// Chargement commande hérité
		for _, precommande := range shell.precommandes {
			// Moins la commande connexion qui n'est pas nécessaire en privilège élevé
			if precommande.Name != "connexion" {
				shelladmin.AddCmd(precommande)
			}
		}
		for _, commande := range shell.commandes {
			shelladmin.AddCmd(commande)
		}
		nvshell = shelladmin
	}

	return nvshell
}
