package platform

import "errors"

type OperatingSystemDescriptor struct {
	Name                string
	Version             *string
	Architecture        string
	Release             *string
	DownloadURLs        []string
	Flavor              *string
	DocumentationURL    *string
	RecommendedHardware Hardware
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
