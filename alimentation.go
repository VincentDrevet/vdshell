package main

import (
	"log"
	"syscall"

	"github.com/abiosoft/ishell"
)

// AjoutPowerCmd Fonction ajoutant la gestion de l'alimentation
func AjoutPowerCmd() *ishell.Cmd {
	powerCmd := &ishell.Cmd{
		Name:     "powerstate",
		Help:     "Permet la gestion de l'alimentation de la carte",
		LongHelp: `Syntaxe : powerstate <option>`,
	}
	powerCmd.AddCmd(&ishell.Cmd{
		Name: "halt",
		Help: "Arrête la carte",
		Func: func(c *ishell.Context) {
			err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
			if err != nil {
				log.Fatalf("%v", err)
			}
		},
	})

	powerCmd.AddCmd(&ishell.Cmd{
		Name: "reboot",
		Help: "Redémarre la carte",
		Func: func(c *ishell.Context) {
			err := syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
			if err != nil {
				log.Fatalf("%v", err)
			}
		},
	})

	return powerCmd
}
