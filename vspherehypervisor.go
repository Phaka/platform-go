package platform

// VSphereHypervisor represents the VMware vSphere hypervisor configuration
type VSphereHypervisor interface {
	// GetGuestOSType returns the guest operating system type
	GetGuestOSType() string
	// GetDiskControllerType returns the disk controller type
	GetDiskControllerType() string
	// GetNetworkAdapterType returns the network adapter type
	GetNetworkAdapterType() string
	// GetFirmware returns the firmware type, e.g. bios, uefi
	GetFirmware() string
	// GetName returns the name of the hypervisor
	GetName() string
	// GetId returns an id of the hypervisor
	GetId() string
}
