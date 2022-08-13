package platform

// Hypervisors represents the hypervisor configurations of an operating system
type Hypervisors interface {
	// GetVMware returns the VMware vSphere hypervisor configuration
	GetVMware() VSphereHypervisor
}
