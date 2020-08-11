package main

import "github.com/abiosoft/ishell"

func main() {

	shelldemarrage := Shell{
		typeshell: Restreint,
		message:   "Acces restreint",
		commandes: []*ishell.Cmd{AjoutShowCommande(), AjoutAuthentification()},
	}

	shellrestreint := NouveauShell(shelldemarrage)

	shellrestreint.Run()

}
