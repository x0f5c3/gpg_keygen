package cmd

import (
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
)

func checkForUpdates() {
	err := pcli.CheckForUpdates()
	pterm.Fatal.PrintOnError(err)

}
