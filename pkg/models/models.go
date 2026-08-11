package models

import "reflect"

type serviceConfiguration struct {
	Image       string            `yaml:"image"`
	Build       string            `yaml:"build"`      // string or object; keep flexible
	DependsOn   []string          `yaml:"depends_on"` // list or map in compose v3+
	Ports       any               `yaml:"ports"`      // list of strings or objects
	Environment map[string]string `yaml:"environment"`
	Restart     string            `yaml:"restart"`
	Healthcheck any               `yaml:"healthcheck"`
	Deploy      any               `yaml:"deploy"`
	// Extra       map[string]any `yaml:",inline"`
}

type DockerComposeConfiguration struct {
	Name     string                          `yaml:"name"`
	Services map[string]serviceConfiguration `yaml:"services"`
	Networks string                          `yaml:"networks"`
	Volumes  string                          `yaml:"volumes"`
	Configs  string                          `yaml:"configs"`
	Secrets  string                          `yaml:"secrets"`
}

func (dc DockerComposeConfiguration) IsEmpty() bool {
	values := reflect.ValueOf(dc)

	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i)

		if value.IsValid() {
			return false
		}
	}

	return true
}
