package platform

type BootMethodsDescriptor struct {
	Http   *BootMethodDescriptor `yaml:"http,omitempty"`
	Cdrom  *BootMethodDescriptor `yaml:"cdrom,omitempty"`
	Floppy *BootMethodDescriptor `yaml:"floppy,omitempty"`
}

func (b *BootMethodsDescriptor) GetHttp() BootMethod {
	return b.Http
}

func (b *BootMethodsDescriptor) GetCdrom() BootMethod {
	return b.Cdrom
}

func (b *BootMethodsDescriptor) GetFloppy() BootMethod {
	return b.Floppy
}
