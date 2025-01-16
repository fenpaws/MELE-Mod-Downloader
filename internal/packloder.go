package internal

import (
	"encoding/json"
	"github.com/fenpaws/MELE-Mod-Downloader/pkg/models"
	"io"
	"os"
)

// LoadAndParseBiq2 loads a JSON file and parses it into a ModCollection struct
func LoadAndParseBiq2(filePath string) (*models.ModCollection, error) {
	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the JSON content into ModCollection struct
	var modCollection models.ModCollection
	err = json.Unmarshal(byteValue, &modCollection)
	if err != nil {
		return nil, err
	}

	return &modCollection, nil
}
