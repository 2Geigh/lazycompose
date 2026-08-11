package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isYamlFilePresent(dirEntries []os.DirEntry) bool {

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		extension := filepath.Ext(entry.Name())
		isYaml := extension == (".yml") || extension == (".yaml")

		if isYaml {
			return true
		}
	}

	return false
}

func main() {

	entries, err := os.ReadDir("./")
	if err != nil {
		panic(fmt.Errorf("failed to read files of current directory: %w", err))
	}

	if !(isYamlFilePresent(entries)) {

		fmt.Println("No Docker Compose configuration file found in this directory.")

		var yesOrNo string
		fmt.Println("Create one? [Y/n]:")
		fmt.Scan(&yesOrNo)

		if strings.ToUpper(yesOrNo) != "Y" {
			return
		}

		const defaultConfigurationName = "docker-compose.yaml"
		_, err := os.Create(defaultConfigurationName)
		if err != nil {
			panic(fmt.Errorf("could not initialize Docker Compose configuration: %w", err))
		}
		configurationName := defaultConfigurationName
		fmt.Printf("Created %s", configurationName)
	}

}
