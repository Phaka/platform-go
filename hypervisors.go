package platform

// Hypervisors represents the hypervisor configurations of an operating system
type Hypervisors interface {
	// GetvSphere returns the VMware vSphere hypervisor configuration
	GetvSphere() VSphereHypervisor
}
