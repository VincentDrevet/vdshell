package main

import (
	"github.com/abiosoft/ishell"
)

func AjoutUserCommande() *ishell.Cmd {
	userCmd := &ishell.Cmd{
		Name:     "user",
		Help:     "Gestion des comptes",
		LongHelp: `Syntaxe : user <element>`,
	}
	passwordSubCmd := &ishell.Cmd{
		Name:     "password",
		Help:     "Gestion des mots de passe",
		LongHelp: "Permet la gestion des mots de passe",
	}

	passwordSubCmd.AddCmd(&ishell.Cmd{
		Name: "reset",
		Help: "Changement du mot de passe",
		Func: func(c *ishell.Context) {
			c.Print("Pour des raisons de sécurité, le mot de passe ne va pas s'afficher à l'écran")
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			var boucle bool = true
			for boucle {
				c.Print("Nouveau mot de passe : ")
				nvmdp := c.ReadPassword()

				c.Print("Confirmation du mot de passe")
				confirmmdp := c.ReadPassword()

				if nvmdp == confirmmdp {
					hash := HashPassword([]byte(nvmdp))
					configuration := ChargerFichierConfiguration("/etc/vdshell/vdshell.ini")
					bdd := ConnexionBDD(configuration.CheminBDD)
					ChangementMotdePasse(bdd, hash)
					c.Print("Mot de passe mis à jour")
					boucle = false
				}

			}
		},
	})

	userCmd.AddCmd(passwordSubCmd)

	return userCmd
}
