package platform

import "fmt"

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
