package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"lazycompose/pkg/models"

	"github.com/goccy/go-yaml"
)

func isContentfulDockerComposeConfigPresent(dirEntries []os.DirEntry) (bool, error) {

	for _, entry := range dirEntries {

		if entry.IsDir() {
			continue
		}

		extension := filepath.Ext(entry.Name())
		isYaml := extension == (".yml") || extension == (".yaml")

		if !isYaml {
			continue
		}

		log.Println(entry.Name())
		log.Println(isYaml)

		content, err := os.ReadFile(entry.Name())
		if err != nil {
			return false, fmt.Errorf("could not open file %s: %w", entry.Name(), err)
		}

		var yaml_contents models.DockerComposeConfiguration
		err = yaml.UnmarshalWithOptions(content, &yaml_contents)
		if err != nil {
			return false, fmt.Errorf("unmarshall %s failed: %w", entry.Name(), err)
		}
		if yaml_contents.IsEmpty() {
			return false, nil
		}

	}

	return false, nil
}

func promptUserToCreateDockerComposeConfigFile() error {
	fmt.Println("No contentful Docker Compose configuration file found in this directory.")

	var yesOrNo string
	fmt.Println("Create one? [Y/n]:")
	fmt.Scan(&yesOrNo)

	if strings.ToUpper(yesOrNo) != "Y" {
		return nil
	}

	err := createDockerComposeConfigFile()
	if err != nil {
		return fmt.Errorf("couldn't create docker compose configruation file: %w", err)
	}

	return nil
}

func createDockerComposeConfigFile() error {
	const defaultConfigurationName = "docker-compose.yaml"
	_, err := os.Create(defaultConfigurationName)
	if err != nil {
		return fmt.Errorf("could not initialize Docker Compose configuration: %w", err)
	}
	configurationName := defaultConfigurationName
	fmt.Printf("Created %s", configurationName)
	return nil
}

func main() {

	entries, err := os.ReadDir("./")
	if err != nil {
		panic(fmt.Errorf("failed to read files of current directory: %w", err))
	}

	isContentfulDockerComposeConfigPresent, err := isContentfulDockerComposeConfigPresent(entries)
	if err != nil {
		panic(fmt.Errorf("failed to determine existence of Docker Compose configuration: %w", err))
	}

	if isContentfulDockerComposeConfigPresent {
		return
	}

	err = promptUserToCreateDockerComposeConfigFile()
	if err != nil {
		panic(fmt.Errorf("prompt user to create Docker Compose configuration file failed: %w", err))
	}

}
