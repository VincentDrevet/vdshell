package main

import "github.com/abiosoft/ishell"

func AjoutAuthentification() *ishell.Cmd {

	authCmd := &ishell.Cmd{
		Name:     "connexion",
		Help:     "Permet de s'authentifier pour accéder au mode privilégié",
		LongHelp: `Syntaxe : connexion`,
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Print("Utilisateur : ")
			utilisateur := c.ReadLine()

			c.Print("Mot de passe : ")
			motdepasse := c.ReadPassword()

			if utilisateur == "admin" && motdepasse == "admin" {
				c.Println("Acces autorisé")
				sh := Shell{
					typeshell:    Administrateur,
					message:      "Entrer dans le mode administrateur",
					precommandes: c.Cmds(),
					commandes:    []*ishell.Cmd{AjoutFSCommande(), AjoutPowerCmd()},
				}
				adminshell := NouveauShell(sh)
				adminshell.Run()

			}

		},
	}

	return authCmd

}
