package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/abiosoft/ishell"
)

// Ajoutsh Fonction ajoutant l'ouverture d'un sh
func Ajoutsh() *ishell.Cmd {
	bashCmd := &ishell.Cmd{
		Name:     "sh",
		Help:     "Lance un shell sh",
		LongHelp: `Syntaxe : sh`,
		Func: func(c *ishell.Context) {
			cmd := exec.Command("/bin/sh")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Fatalf("%v", err)
			}
		},
	}

	return bashCmd

}
