package media

import (
	"encoding/json"
	"fmt"
	"goth-todo/internal/models"
	"io"
	"os"
	"path/filepath"
)

var manifest models.Manifest

func GetCSSHashFile() string {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		os.Exit(1)
	}

	file := "public/manifest.json"

	manifestPath, err := filepath.Abs(filepath.Join(wd, file))
	if err != nil {
		fmt.Println("Error resolving manifest path:", err)
		os.Exit(1)
	}

	// Check if the file exists before proceeding
	if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
		fmt.Println("Error: manifest file not found at", manifestPath)
		os.Exit(1)
	}

	// Save the hashed filename in a JSON manifest
	fmt.Printf("Getting file from %s", manifestPath)

	jsonData, err := os.Open(manifestPath)

	if err != nil {
		fmt.Println("Error opening manifest:", err)
		os.Exit(1)
	}

	defer jsonData.Close()

	manifestBytes, _ := io.ReadAll(jsonData)

	json.Unmarshal(manifestBytes, &manifest)

	return fmt.Sprintf("/public/styles/%s", manifest.CSS)
}
