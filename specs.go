package platform

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type validatable interface {
	validate() error
}

type Processors struct {
	Count int `yaml:"count,omitempty"`
	Cores int `yaml:"cores,omitempty"`
}

func (p *Processors) validate() error {
	if p.Count <= 0 {
		return fmt.Errorf("platform: processors.count must be greater than 0")
	}
	if p.Cores <= 0 {
		return fmt.Errorf("platform: processors.cores must be greater than 0")
	}
	return nil
}

func (p *Processors) String() string {
	return toYAML(p)
}

type Hardware struct {
	Memory     int        `yaml:"memory,omitempty"`
	Storage    int        `yaml:"storage,omitempty"`
	Processors Processors `yaml:"processors,omitempty"`
}

func (h *Hardware) validate() error {
	if h.Memory <= 0 {
		return fmt.Errorf("platform: hardware.memory must be greater than 0")
	}
	if h.Storage <= 0 {
		return fmt.Errorf("platform: hardware.storage must be greater than 0")
	}
	return h.Processors.validate()
}

func (h *Hardware) String() string {
	return toYAML(h)
}

type Hypervisor map[string]interface{}

func (h Hypervisor) String() string {
	return toYAML(h)
}

type HypervisorKind string

const (
	HypervisorKindVMware = "vSphere"
)

type Hypervisors map[string]Hypervisor

func (h Hypervisors) validate() error {
	if h == nil {
		return fmt.Errorf("platform: hypervisors is required")
	}

	if len(h) != 1 {
		return fmt.Errorf("platform: only one hypervisor is supported")
	}

	for _, kind := range AllHypervisorKinds() {
		if _, ok := h[kind]; !ok {
			return fmt.Errorf("platform: hypervisors \"%s\" is required", kind)
		}
	}

	return nil
}

func (h Hypervisors) String() string {
	return toYAML(h)
}

type Variable struct {
	Name         string `yaml:"name,omitempty"`
	DefaultValue string `yaml:"default,omitempty"`
	Type         string `yaml:"type,omitempty"`
}

func (v *Variable) String() string {
	return toYAML(v)
}

type BootMethod struct {
	Name      string            `yaml:"name,omitempty"`
	Commands  string            `yaml:"commands,omitempty"`
	Files     map[string]string `yaml:"files,omitempty"`
	Delay     string            `yaml:"delay,omitempty"`
	Variables []*Variable       `yaml:"variables,omitempty"`
}

type BootMethodKind string

const (
	BootMethodKindHttp   = "http"
	BootMethodKindCdrom  = "cdrom"
	BootMethodKindFloppy = "floppy"
)

func (b *BootMethod) String() string {
	return toYAML(b)
}

type BootMethods map[string]*BootMethod

func (b BootMethods) String() string {
	return toYAML(b)
}

func (b BootMethods) validate() error {
	if b == nil {
		return fmt.Errorf("platform: boot methods is required")
	}

	if len(b) == 0 {
		return fmt.Errorf("platform: at least one boot method is required")
	}

	for _, kind := range BootMethodKinds() {
		if _, ok := b[kind]; !ok {
			return fmt.Errorf("platform: boot methods \"%s\" is required", kind)
		}
	}
	return nil

}

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

var defaultHardware = Hardware{
	Memory:  2048,
	Storage: 8192,
	Processors: Processors{
		Count: 1,
		Cores: 1,
	},
}

var defaultHypervisors = Hypervisors{
	HypervisorKindVMware: Hypervisor{
		"Name":               "VMware vSphere",
		"Id":                 "vsphere",
		"GuestOSType":        "otherGuest64",
		"Firmware":           "bios",
		"NetworkAdapterType": "e1000",
		"DiskControllerType": "lsilogic",
	},
}

var defaultBootMethods = BootMethods{
	BootMethodKindHttp: &BootMethod{
		Name:     "HTTP",
		Commands: "",
		Files:    map[string]string{},
	},
	BootMethodKindCdrom: &BootMethod{
		Name:     "CD-ROM",
		Commands: "",
		Files:    map[string]string{},
	},
	BootMethodKindFloppy: &BootMethod{
		Name:     "Floppy",
		Commands: "",
		Files:    map[string]string{},
	},
}

var defaultOperatingSystem = OperatingSystem{
	Name:                "",
	Id:                  "",
	Version:             "",
	Architecture:        "amd64",
	Release:             "",
	DownloadURLs:        []string{},
	Flavor:              "",
	DocumentationURL:    "",
	RecommendedHardware: defaultHardware,
	Hypervisors:         defaultHypervisors,
	BootMethods:         defaultBootMethods,
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

			switch kind {
			case HypervisorKindVMware:
				hypervisor, _ := o.Hypervisors[kind]
				oldNames := map[string]string{
					"guest_os_type":        "GuestOSType",
					"disk_controller_type": "DiskControllerType",
					"network_adapter_type": "NetworkAdapterType",
					"firmware":             "Firmware",
				}
				for old, new := range oldNames {
					hypervisor[new] = hypervisor[old]
					delete(hypervisor, old)
				}
			}
		}
	}

	if o.BootMethods == nil {
		o.BootMethods = defaultBootMethods
	} else {

		for _, kind := range BootMethodKinds() {
			kindLower := strings.ToLower(kind)
			if _, ok := o.BootMethods[kindLower]; ok {
				// so we have a lowercase key, but we really need to case it
				// to match the kind
				o.BootMethods[kind] = o.BootMethods[kindLower]
				delete(o.BootMethods, kindLower)
			}

			if _, ok := o.BootMethods[kind]; !ok {
				o.BootMethods[kind] = defaultBootMethods[kind]
			}
		}
	}

	return o.validate()
}

func AllHypervisorKinds() []string {
	return []string{
		HypervisorKindVMware,
	}
}

func BootMethodKinds() []string {
	return []string{
		BootMethodKindHttp,
		BootMethodKindCdrom,
		BootMethodKindFloppy,
	}
}

type OperatingSystems []*OperatingSystem

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

// toString returns the string YAML representation of the given value.
func toYAML(p interface{}) string {
	bytes, err := yaml.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}
