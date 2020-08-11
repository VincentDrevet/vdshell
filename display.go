package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func dessinertableau(titres []string, donnees [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(titres)
	table.SetBorder(true)     // Set Border to false
	table.AppendBulk(donnees) // Add Bulk Data
	table.Render()
}
