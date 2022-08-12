package platform

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// OperatingSystem represents the configuration of an operating system
type OperatingSystem interface {
	// GetName returns the name of the operating system
	GetName() string
	// GetVersion returns the version of the operating system, e.g. 20.04.3
	GetVersion() string
	// GetArchitecture returns the architecture of the operating system, e.g. amd64
	GetArchitecture() string
	// GetRelease returns the release of the operating system, e.g. bionic
	GetRelease() string
	// GetDownloadURLs returns the download URLs of the operating system
	GetDownloadURLs() []string
	// GetFlavor returns the flavor of the operating system, e.g. desktop or server
	GetFlavor() string
	// GetDocumentationURL returns the URL to the operating system documentation
	GetDocumentationURL() string
	// GetRecommendedHardware returns the recommended hardware of the operating system
	GetRecommendedHardware() Hardware
	// Validate returns an error if the operating system configuration is valid
	Validate() error
	// GetHypervisors returns the hypervisor configurations of the operating system
	GetHypervisors() Hypervisors
	// GetBootMethods returns the methods of booting and installing the operating system
	GetBootMethods() BootMethods
}

type OperatingSystems []OperatingSystem

// LoadOperatingSystem loads the operating system configuration from the given YAML file
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
