package platform

type BootMethodsDescriptor struct {
	Http   *BootMethodDescriptor `yaml:"http"`
	Cdrom  *BootMethodDescriptor `yaml:"cdrom"`
	Floppy *BootMethodDescriptor `yaml:"floppy"`
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
