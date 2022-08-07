package platform

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type OperatingSystem interface {
	GetName() string
	GetVersion() string
	GetArchitecture() string
	GetRelease() string
	GetDownloadURLs() []string
	GetFlavor() string
	GetDocumentationURL() string
	GetRecommendedHardware() Hardware
	Validate() error
}

type OperatingSystems []OperatingSystem

func LoadOperatingSystem(path string) (OperatingSystem, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	desc := &OperatingSystemDescriptor{}
	err = yaml.Unmarshal(yamlFile, desc)
	if err != nil {
		return nil, err
	}
	return desc, nil
}
