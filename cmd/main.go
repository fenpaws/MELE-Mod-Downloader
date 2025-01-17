package main

import (
	"github.com/fenpaws/MELE-Mod-Downloader/cmd/cmd"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/utils"
)

func init() {
	// Set up the logger based on the configuration
	utils.SetupLogger("INFO", "PLAIN")

}

func main() {
	cmd.Execute()
}
