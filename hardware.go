package platform

import "fmt"

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
