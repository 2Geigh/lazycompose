package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isDockerComposeConfigPresent(dirEntries []os.DirEntry) (bool, error) {

	type dockerConfiguration struct {
		name     string `yaml:"name"`
		services string `yaml:"services"`
		networks string `yaml:"networks"`
		volumes  string `yaml:"volumes"`
		configs  string `yaml:"configs"`
		secrets  string `yaml:"secrets"`
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}

		extension := filepath.Ext(entry.Name())
		isYaml := extension == (".yml") || extension == (".yaml")

		if isYaml {
			return true, nil
		}

		// file, err := os.ReadFile(entry.Name())
		// if err != nil {
		// 	return false, fmt.Errorf("could not read file %s: %w", entry.Name(), err)
		// }

		// var yaml_contents dockerConfiguration
		// err = yaml.Unmarshal([]byte(file), &yaml_contents)
		// if err != nil {
		// 	return false, fmt.Errorf("unmarshall %s failed: %w", entry.Name(), err)
		// }

	}

	return false, nil
}

func main() {

	entries, err := os.ReadDir("./")
	if err != nil {
		panic(fmt.Errorf("failed to read files of current directory: %w", err))
	}

	isDockerComposeConfigPresent, err := isDockerComposeConfigPresent(entries)
	if err != nil {
		panic(fmt.Errorf("failed to determine existence of Docker Compose configuration: %w", err))
	}

	if !(isDockerComposeConfigPresent) {

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
