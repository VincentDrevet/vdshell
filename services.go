package main

import (
	"log"
	"os/exec"

	"github.com/abiosoft/ishell"
)

// AddServiceCmd ajoute la gestion des services
func AddServiceCmd() *ishell.Cmd {
	serviceCmd := &ishell.Cmd{
		Name:     "service",
		Help:     "Permet la gestion des services",
		LongHelp: `Syntaxe : service <service>`,
	}

	sshservice := &ishell.Cmd{
		Name:     "ssh",
		Help:     "serveur ssh",
		LongHelp: "Gestion du service SSH",
	}

	sshservice.AddCmd(&ishell.Cmd{
		Name: "start",
		Help: "Démarre le service SSH",
		Func: func(c *ishell.Context) {
			cmd := exec.Command("/etc/init.d/S50sshd", "start")
			output, err := cmd.Output()
			if err != nil {
				log.Fatalf("%v", output)
			}
			c.Printf("%s", output)
		},
	})

	sshservice.AddCmd(&ishell.Cmd{
		Name: "stop",
		Help: "Arrête le service SSH",
		Func: func(c *ishell.Context) {
			cmd := exec.Command("/etc/init.d/S50sshd", "stop")
			output, err := cmd.Output()
			if err != nil {
				log.Fatalf("%v", output)
			}
			c.Printf("%s", output)
		},
	})

	serviceCmd.AddCmd(sshservice)

	return serviceCmd
}
