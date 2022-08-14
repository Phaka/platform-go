package platform

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Parse parses the given YAML bytes into an OperatingSystem.
func Parse(b []byte) (*OperatingSystem, error) {
	if len(b) == 0 {
		return nil, nil
	}
	var o = &OperatingSystem{}
	err := yaml.Unmarshal(b, o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// ParseFile parses the given YAML file into an OperatingSystem.
func ParseFile(path string) (*OperatingSystem, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	d := &OperatingSystem{}
	d, err = Parse(b)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// ParseGlob parses all files matching the given glob pattern into a slice of OperatingSystems.
func ParseGlob(pattern string) (result OperatingSystems, err error) {
	filenames, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("platform: pattern matches no files: %#q", pattern)
	}
	return ParseFiles(filenames...)
}

// ParseFiles parses the given YAML files into a slice of OperatingSystems.
func ParseFiles(filenames ...string) (OperatingSystems, error) {
	var result OperatingSystems
	for _, filename := range filenames {
		desc, err := ParseFile(filename)
		if err != nil {
			return nil, err
		}
		result = append(result, desc)
	}

	return result, nil
}
