package platform

type VariableDescriptor struct {
	Name         string  `yaml:"name,omitempty"`
	DefaultValue *string `yaml:"default,omitempty"`
	Type         *string `yaml:"type,omitempty"`
}

func (v *VariableDescriptor) GetName() string {
	return v.Name
}

func (v *VariableDescriptor) SetName(value string) {
	v.Name = value
}

func (v *VariableDescriptor) GetDefaultValue() *string {
	return v.DefaultValue
}

func (v *VariableDescriptor) SetDefaultValue(value string) {
	if value != "" {
		v.DefaultValue = &value
	} else {
		v.DefaultValue = nil
	}
}

func (v *VariableDescriptor) GetType() string {
	if v.Type == nil {
		return "string"
	}
	return *v.Type
}

func (v *VariableDescriptor) SetType(value string) {
	if value != "" {
		v.Type = &value
	} else {
		v.Type = nil
	}
}
