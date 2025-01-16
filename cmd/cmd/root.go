package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gole",
	Short: "Mass Effect: Legendary Edition mod downloader for Nexus Mods",
}

func init() {
	rootCmd.AddCommand(consumeCmd)
	rootCmd.AddCommand(runCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Error("could not execute root command")
		os.Exit(1)
	}
}
