package platform

type BootMethodDescriptor struct {
	Commands  *string               `yaml:"commands,omitempty"`
	Files     map[string]string     `yaml:"files,omitempty"`
	Delay     *string               `yaml:"delay,omitempty"`
	Variables []*VariableDescriptor `yaml:"variables,omitempty"`
}

func (b *BootMethodDescriptor) GetWait() string {
	if b.Delay == nil {
		return "5s"
	}
	return *b.Delay
}

func (b *BootMethodDescriptor) GetCommands() string {
	if b.Commands == nil {
		return ""
	}
	return *b.Commands
}

func (b *BootMethodDescriptor) GetFiles() map[string]string {
	return b.Files
}

func (b *BootMethodDescriptor) GetVariables() []Variable {
	variables := make([]Variable, 0, len(b.Variables))
	for _, v := range b.Variables {
		variables = append(variables, v)
	}
	return variables
}
