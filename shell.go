package main

import "github.com/abiosoft/ishell"

type TypeShell int

const (
	Restreint TypeShell = iota
	Administrateur
)

type Shell struct {
	typeshell    TypeShell
	message      string
	commandes    []*ishell.Cmd
	precommandes []*ishell.Cmd
}

func NouveauShell(shell Shell) *ishell.Shell {

	var nvshell *ishell.Shell
	switch shell.typeshell {
	case Restreint:
		sh := ishell.New()
		sh.Println(shell.message)
		sh.SetPrompt("Mode restreint> ")
		for _, commande := range shell.commandes {
			sh.AddCmd(commande)
		}
		nvshell = sh

	case Administrateur:
		shelladmin := ishell.New()
		shelladmin.Println(shell.message)
		shelladmin.SetPrompt("Mode administrateur> ")
		// Chargement commande hérité
		for _, precommande := range shell.precommandes {
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
