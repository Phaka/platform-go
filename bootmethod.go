package platform

import (
	"fmt"
	"time"
)

type BootMethod struct {
	Name      string            `yaml:"name,omitempty"`
	Commands  string            `yaml:"commands,omitempty"`
	Files     map[string]string `yaml:"files,omitempty"`
	Delay     string            `yaml:"delay,omitempty"`
	Variables []*Variable       `yaml:"variables,omitempty"`
	Username  string            `yaml:"username,omitempty"`
}

func (b *BootMethod) validate() error {
	if b == nil {
		return nil
	}
	if b.Name == "" {
		return fmt.Errorf("platform: boot method name is required")
	}
	if b.Delay == "" {
		return fmt.Errorf("platform: boot method delay is required")
	}
	_, err := time.ParseDuration(b.Delay)
	if err != nil {
		return fmt.Errorf("platform: boot method delay is invalid: %s", err)
	}

	return nil
}

func (b *BootMethod) String() string {
	return toYAML(b)
}
