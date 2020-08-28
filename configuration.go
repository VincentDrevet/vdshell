package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Configuration struct {
	CheminBDD string
}

// FichierConfigurationExiste vérifie l'existance du fichier de configuration
func FichierConfigurationExiste(chemin string) bool {
	var retour bool
	if _, err := os.Stat(chemin); os.IsNotExist(err) {
		fmt.Printf("%v", err)
		retour = false
	} else {
		retour = true
	}
	return retour
}

func ChargerFichierConfiguration(chemin string) Configuration {

	cfg, err := ini.Load(chemin)
	if err != nil {
		fmt.Printf("%v", err)
	}
	configuration := Configuration{CheminBDD: cfg.Section("BDD").Key("chemin").String()}
	return configuration
}
