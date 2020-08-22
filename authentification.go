package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/abiosoft/ishell"
)

// AjoutAuthentification Ajoutant la gestion de l'authentification pour l'acces administrateur
func AjoutAuthentification() *ishell.Cmd {

	authCmd := &ishell.Cmd{
		Name:     "enable",
		Help:     "Permet de s'authentifier pour accéder au mode privilégié",
		LongHelp: `Syntaxe : enable`,
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)

			c.Print("Utilisateur : ")
			utilisateur := c.ReadLine()

			c.Print("Mot de passe : ")
			motdepasse := c.ReadPassword()

			bdd := ConnexionBDD("./central.db")
			mdpbase := Recuperationmotdepasse(bdd, utilisateur)

			// si le mot de passe est vide alors l'utilisateur n'existe pas
			if mdpbase == "" {
				fmt.Println("Utilisateur inconnu")
				return
			}

			// On vérifie le hash en base avec le mot de passe fournie par l'utilisateur
			if bcrypt.CompareHashAndPassword([]byte(mdpbase), []byte(motdepasse)) == nil {
				c.Println("Acces autorisé")
				sh := Shell{
					typeshell:    Administrateur,
					message:      "Entrer dans le mode administrateur",
					precommandes: c.Cmds(),
					commandes:    []*ishell.Cmd{AjoutFSCommande(), AjoutPowerCmd(), AddServiceCmd(), Ajoutsh(), AjoutUserCommande()},
				}
				bdd.Close()
				adminshell := NouveauShell(sh)
				adminshell.Run()
			}
		},
	}

	return authCmd

}
