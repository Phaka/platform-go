package platform

import (
	"log"

	"gopkg.in/yaml.v3"
)

// toString returns the string YAML representation of the given value.
func toYAML(p interface{}) string {
	bytes, err := yaml.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}
