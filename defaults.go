package platform

var defaultHardware = Hardware{
	Memory:  2048,
	Storage: 8192,
	Processors: Processors{
		Count: 1,
		Cores: 1,
	},
}

var defaultHypervisors = Hypervisors{
	HypervisorKindVMware: &Hypervisor{
		"Name":               "VMware vSphere",
		"Id":                 "vsphere",
		"GuestOSType":        "otherGuest64",
		"Firmware":           "bios",
		"NetworkAdapterType": "e1000",
		"DiskControllerType": "lsilogic",
	},
}

var defaultBootMethods = BootMethods{
	BootMethodKindHttp: &BootMethod{
		Name:      "HTTP",
		Commands:  "",
		Files:     map[string]string{},
		Delay:     "5s",
		Variables: []*Variable{},
		Username:  "root",
	},
	BootMethodKindCdrom: &BootMethod{
		Name:      "CD-ROM",
		Commands:  "",
		Files:     map[string]string{},
		Delay:     "5s",
		Variables: []*Variable{},
		Username:  "root",
	},
	BootMethodKindFloppy: &BootMethod{
		Name:      "Floppy",
		Commands:  "",
		Files:     map[string]string{},
		Delay:     "5s",
		Variables: []*Variable{},
		Username:  "root",
	},
}

var defaultOperatingSystem = OperatingSystem{
	Name:                "",
	Id:                  "",
	Version:             "",
	Architecture:        "amd64",
	Release:             "",
	DownloadURLs:        []string{},
	Flavor:              "",
	DocumentationURL:    "",
	RecommendedHardware: defaultHardware,
	Hypervisors:         defaultHypervisors,
	BootMethods:         defaultBootMethods,
}
