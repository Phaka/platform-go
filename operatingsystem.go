package platform

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type OperatingSystem struct {
	Name                string      `yaml:"name"`
	Id                  string      `yaml:"id"`
	Version             string      `yaml:"version"`
	Architecture        string      `yaml:"architecture"`
	Release             string      `yaml:"release"`
	DownloadURLs        []string    `yaml:"downloads"`
	Flavor              string      `yaml:"flavor"`
	DocumentationURL    string      `yaml:"documentation"`
	RecommendedHardware Hardware    `yaml:"hardware"`
	Hypervisors         Hypervisors `yaml:"hypervisors"`
	BootMethods         BootMethods `yaml:"boot"`
}

func (o *OperatingSystem) validate() (err error) {
	if o.Name == "" {
		return fmt.Errorf("platform: operating system name is required")
	}
	err = o.RecommendedHardware.validate()
	if err != nil {
		return err
	}
	err = o.Hypervisors.validate()
	if err != nil {
		return err
	}
	err = o.BootMethods.validate()
	if err != nil {
		return err
	}
	return nil
}

func (o *OperatingSystem) String() string {
	return toYAML(o)
}

func (o *OperatingSystem) UnmarshalYAML(value *yaml.Node) error {

	type plain OperatingSystem
	if err := value.Decode((*plain)(o)); err != nil {
		return err
	}

	o.Name = strings.TrimSpace(o.Name)
	o.Id = strings.TrimSpace(o.Id)
	o.Version = strings.TrimSpace(o.Version)
	o.Architecture = strings.TrimSpace(o.Architecture)
	o.Release = strings.TrimSpace(o.Release)
	o.Flavor = strings.TrimSpace(o.Flavor)
	o.DocumentationURL = strings.TrimSpace(o.DocumentationURL)

	if o.Id == "" {
		o.Id = strings.ToLower(o.Name)
	}

	if o.Architecture == "" {
		o.Architecture = defaultOperatingSystem.Architecture
	}

	if o.DownloadURLs == nil {
		o.DownloadURLs = defaultOperatingSystem.DownloadURLs
	}

	if o.RecommendedHardware.Memory <= 0 {
		o.RecommendedHardware.Memory = defaultHardware.Memory
	}

	if o.RecommendedHardware.Storage <= 0 {
		o.RecommendedHardware.Storage = defaultHardware.Storage
	}

	if o.RecommendedHardware.Processors.Count <= 0 {
		o.RecommendedHardware.Processors.Count = defaultHardware.Processors.Count
	}

	if o.RecommendedHardware.Processors.Cores <= 0 {
		o.RecommendedHardware.Processors.Cores = defaultHardware.Processors.Cores
	}

	if o.Hypervisors == nil {
		o.Hypervisors = defaultHypervisors
	} else {
		for _, kind := range AllHypervisorKinds() {

			kindLower := strings.ToLower(kind)
			if _, ok := o.Hypervisors[kindLower]; ok {
				// so we have a lowercase key, but we really need to case it
				// to match the kind
				o.Hypervisors[kind] = o.Hypervisors[kindLower]
				delete(o.Hypervisors, kindLower)

			}

			if _, ok := o.Hypervisors[kind]; !ok {
				o.Hypervisors[kind] = defaultHypervisors[kind]
			}
		}
	}

	if o.BootMethods == nil {
		o.BootMethods = defaultBootMethods
	} else {

		for _, kind := range BootMethodKinds() {
			// kindLower := strings.ToLower(kind)
			// if _, ok := o.BootMethods[kindLower]; ok {
			// 	// so we have a lowercase key, but we really need to case it
			// 	// to match the kind
			// 	o.BootMethods[kind] = o.BootMethods[kindLower]
			// 	delete(o.BootMethods, kindLower)
			// }

			if _, ok := o.BootMethods[kind]; !ok {
				// o.BootMethods[kind] = nil
			} else {
				// check if boot method name exists, if not assign it
				if o.BootMethods[kind].Name == "" {
					o.BootMethods[kind].Name = defaultBootMethods[kind].Name
				}
				if o.BootMethods[kind].Username == "" {
					o.BootMethods[kind].Username = defaultBootMethods[kind].Username
				}
			}

		}
	}

	return o.validate()
}
