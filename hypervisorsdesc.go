package platform

type HypervisorsDescriptor struct {
	VSphere *VSphereHypervisorDescriptor `yaml:"vsphere,omitempty"`
}

func (h *HypervisorsDescriptor) GetVMware() VSphereHypervisor {
	return h.VSphere
}
