package platform

type Variable struct {
	Name         string `yaml:"name,omitempty"`
	DefaultValue string `yaml:"default,omitempty"`
	Type         string `yaml:"type,omitempty"`
}

func (v *Variable) String() string {
	return toYAML(v)
}
