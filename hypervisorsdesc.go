package platform

type HypervisorsDescriptor struct {
    VSphere *VSphereHypervisorDescriptor `yaml:"vsphere,omitempty"`
}

func (h *HypervisorsDescriptor) GetvSphere() VSphereHypervisor {
    return h.VSphere
}
