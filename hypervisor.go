package platform

import "fmt"

type Hypervisor map[string]interface{}

func (h Hypervisor) String() string {
	return toYAML(h)
}

func (h Hypervisor) validate() error {
	if h == nil {
		return fmt.Errorf("platform: hypervisor is required")
	}

	return nil
}

func (h Hypervisor) Name() string {
	// check if name exists in h and return it
	if name, ok := h["name"]; ok {
		return name.(string)
	}
	if name, ok := h["Name"]; ok {
		return name.(string)
	}
	return ""
}

func (h Hypervisor) SetName(value string) {
	h["name"] = value
}

func (h Hypervisor) Id() string {
	// check if name exists in h and return it
	if name, ok := h["id"]; ok {
		return name.(string)
	}
	if name, ok := h["Id"]; ok {
		return name.(string)
	}
	return ""
}

func (h Hypervisor) SetId(value string) {
	h["id"] = value
}

func (h Hypervisor) GuestOSType() string {
	// check if guestOSType exists in h and return it
	if guestOSType, ok := h["guest_os_type"]; ok {
		return guestOSType.(string)
	}
	if guestOSType, ok := h["GuestOSType"]; ok {
		return guestOSType.(string)
	}
	return ""
}

func (h Hypervisor) SetGuestOSType(value string) {
	h["guest_os_type"] = value
}

func (h Hypervisor) Firmware() string {
	// check if firmware exists in h and return it
	if firmware, ok := h["firmware"]; ok {
		return firmware.(string)
	}
	if firmware, ok := h["Firmware"]; ok {
		return firmware.(string)
	}
	return ""
}

func (h Hypervisor) SetFirmware(value string) {
	h["firmware"] = value
}

func (h Hypervisor) NetworkAdapterType() string {
	// check if networkAdapterType exists in h and return it
	if networkAdapterType, ok := h["network_adapter_type"]; ok {
		return networkAdapterType.(string)
	}
	if networkAdapterType, ok := h["NetworkAdapterType"]; ok {
		return networkAdapterType.(string)
	}
	return ""
}

func (h Hypervisor) SetNetworkAdapterType(value string) {
	h["network_adapter_type"] = value
}

func (h Hypervisor) DiskControllerType() string {
	// check if diskControllerType exists in h and return it
	if diskControllerType, ok := h["disk_controller_type"]; ok {
		return diskControllerType.(string)
	}
	if diskControllerType, ok := h["DiskControllerType"]; ok {
		return diskControllerType.(string)
	}
	return ""
}

func (h Hypervisor) SetDiskControllerType(value string) {
	h["disk_controller_type"] = value
}
