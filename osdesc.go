package platform

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type OperatingSystemDescriptor struct {
	Name                string                 `yaml:"name,omitempty"`
	Version             *string                `yaml:"version,omitempty"`
	Architecture        string                 `yaml:"architecture,omitempty"`
	Release             *string                `yaml:"release,omitempty"`
	DownloadURLs        []string               `yaml:"downloads,omitempty"`
	Flavor              *string                `yaml:"flavor,omitempty"`
	DocumentationURL    *string                `yaml:"documentation,omitempty"`
	RecommendedHardware *HardwareDescriptor    `yaml:"hardware,omitempty"`
	BootMethods         *BootMethodsDescriptor `yaml:"boot,omitempty"`
	Hypervisors         *HypervisorsDescriptor `yaml:"hypervisors,omitempty"`
}

func (o *OperatingSystemDescriptor) Save(path string) error {
	bytes, err := yaml.Marshal(o)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0644)
}

var defaultBootMethod = &BootMethodDescriptor{
	Commands: nil,
	Files:    nil,
}

var defaultBootMethods = &BootMethodsDescriptor{
	Http:   defaultBootMethod,
	Cdrom:  defaultBootMethod,
	Floppy: defaultBootMethod,
}

func (o *OperatingSystemDescriptor) GetBootMethods() BootMethods {
	if o.BootMethods == nil {
		return defaultBootMethods
	}
	return o.BootMethods
}

var defaultHypervisors = &HypervisorsDescriptor{
	VSphere: &VSphereHypervisorDescriptor{
		GuestOSType:        nil,
		DiskControllerType: nil,
		NetworkAdapterType: nil,
		Firmware:           nil,
	},
}

func (o *OperatingSystemDescriptor) Parse(b []byte) error {
	err := yaml.Unmarshal(b, o)
	if err != nil {
		return err
	}
	return nil
}

func (o *OperatingSystemDescriptor) GetHypervisors() Hypervisors {
	if o.Hypervisors == nil {
		return defaultHypervisors
	}
	return o.Hypervisors
}

func (o *OperatingSystemDescriptor) String() string {
	bytes, err := yaml.Marshal(o)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (o *OperatingSystemDescriptor) Validate() error {
	if o.Name == "" {
		return errors.New("invalid os name")
	}
	if o.Architecture == "" {
		return errors.New("invalid architecture")
	}
	if o.RecommendedHardware != nil {
		return o.RecommendedHardware.Validate()
	}
	return nil
}

func (o *OperatingSystemDescriptor) GetName() string {
	return o.Name
}

func (o *OperatingSystemDescriptor) GetVersion() string {
	if o.Version == nil {
		return ""
	}
	return *o.Version
}

func (o *OperatingSystemDescriptor) GetArchitecture() string {
	return o.Architecture
}

func (o *OperatingSystemDescriptor) GetRelease() string {
	if o.Release == nil {
		return ""
	}
	return *o.Release
}

func (o *OperatingSystemDescriptor) GetDownloadURLs() []string {
	return o.DownloadURLs
}

func (o *OperatingSystemDescriptor) GetFlavor() string {
	if o.Flavor == nil {
		return ""
	}
	return *o.Flavor
}

func (o *OperatingSystemDescriptor) GetDocumentationURL() string {
	if o.DocumentationURL == nil {
		return ""
	}
	return *o.DocumentationURL
}

var defaultHardware = &HardwareDescriptor{
	Memory:     nil,
	Storage:    nil,
	Processors: nil,
}

func (o *OperatingSystemDescriptor) GetRecommendedHardware() Hardware {
	if o.RecommendedHardware == nil {
		return defaultHardware
	}
	return o.RecommendedHardware
}
