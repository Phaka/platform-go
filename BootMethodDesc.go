package platform

type BootMethodDescriptor struct {
	Commands *string           `yaml:"commands,omitempty"`
	Files    map[string]string `yaml:"files,omitempty"`
	Delay    *string           `yaml:"delay,omitempty"`
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
