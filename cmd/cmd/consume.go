package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/api/v1/consume"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
)

var consumeCmd = &cobra.Command{
	Use:   "consume [nxm://]",
	Short: "Gets the callback URL from NexusMods",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sendRequest(args[0])
	},
}

func sendRequest(nexusResponse string) {
	// Define the URL for the POST request
	url := "http://localhost:8081/api/v1/consume"

	// Create the JSON body
	data := consume.NexusRequest{NexusURL: nexusResponse}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	// Send the POST request with the JSON body
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error making POST request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Print the response from the server
	log.Printf("Response status code: %d", resp.StatusCode)
}
