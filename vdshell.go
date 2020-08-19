package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/abiosoft/ishell"
)

func main() {

	// On définie la taille du terminal pour éviter des problèmes notamment en console
	cmd := exec.Command("stty", "rows", "60", "cols", "160")
	cmd.Stdin = os.Stdin
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("%v", err)
	}

	shelldemarrage := Shell{
		typeshell: Restreint,
		message:   "Acces restreint",
		commandes: []*ishell.Cmd{AjoutShowCommande(), AjoutAuthentification()},
	}

	shellrestreint := NouveauShell(shelldemarrage)

	shellrestreint.Run()

}
