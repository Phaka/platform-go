package platform

import "errors"

type ProcessorsDescriptor struct {
	Count *int `yaml:"count,omitempty"`
	Cores *int `yaml:"cores,omitempty"`
}

func (p *ProcessorsDescriptor) GetCount() int {
	if p.Count == nil {
		return 1
	}
	return *p.Count
}

func (p *ProcessorsDescriptor) GetCores() int {
	if p.Cores == nil {
		return 1
	}
	return *p.Cores
}

func (p *ProcessorsDescriptor) Validate() error {
	if p.Count != nil && *p.Count < 0 {
		return errors.New("invalid processor count")
	}
	if p.Cores != nil && *p.Cores < 0 {
		return errors.New("invalid processor core count")
	}
	return nil
}
