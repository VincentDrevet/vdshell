package main

import (
	"syscall"

	"github.com/abiosoft/ishell"
)

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
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
		},
	})

	powerCmd.AddCmd(&ishell.Cmd{
		Name: "reboot",
		Help: "Redémarre la carte",
		Func: func(c *ishell.Context) {
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART2)
		},
	})

	return powerCmd

}
