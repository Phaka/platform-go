package platform

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
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

	// GetHypervisors returns the hypervisor configurations of the operating system
	GetHypervisors() Hypervisors

	// GetBootMethods returns the methods of booting and installing the operating system
	GetBootMethods() BootMethods

	// Validate returns an error if the operating system configuration is valid
	Validate() error

	// Save writes the operating system configuration to the given YAML file
	Save(path string) error
}

type OperatingSystems []OperatingSystem

// LoadOperatingSystem loads the operating system configuration from the given YAML file
func LoadOperatingSystem(path string) (OperatingSystem, error) {
	return ParseFile(path)
}

func ParseFile(path string) (OperatingSystem, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	d := &OperatingSystemDescriptor{}
	err = d.Parse(b)
	if err != nil {
		return nil, err
	}
	return d, nil
}

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

func ParseFiles(filenames ...string) (OperatingSystems, error) {
	var result OperatingSystems
	for _, filename := range filenames {
		desc, err := LoadOperatingSystem(filename)
		if err != nil {
			return nil, err
		}
		result = append(result, desc)
	}

	return result, nil
}
