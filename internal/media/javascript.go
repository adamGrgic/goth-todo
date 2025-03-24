package media

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetHTMX() string {
	return "../../../public/scripts/htmx/htmx.min.js"
}

func GetHTMXResponseTargets() string {
	return "../../../public/scripts/htmx/response-targets.js"
}

// todo: create function that can get an array of modules
func GetJsFile(file string) string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		os.Exit(1)
	}

	absolutePath, err := filepath.Abs(filepath.Join(wd, "/public/manifest.json"))
	if err != nil {
		fmt.Println("Error resolving manifest path:", err)
		os.Exit(1)
	}

	// Check if the file exists before proceeding
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		fmt.Println("Error: manifest file not found at", absolutePath)
		os.Exit(1)
	}

	// Open the file
	manifest, err := os.Open(absolutePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer manifest.Close()

	var data map[string]string

	// Decode the JSON into the map
	decoder := json.NewDecoder(manifest)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	fmt.Println(data)

	return data[file]
}

// given component/module names, retrieve ... something ... that can load all those scripts
// func GetJSFiles() string {

// }
