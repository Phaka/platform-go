package platform

import "gopkg.in/yaml.v3"

type VSphereHypervisorDescriptor struct {
    GuestOSType        *string `yaml:"guest_os_type,omitempty"`
    DiskControllerType *string `yaml:"disk_controller_type,omitempty"`
    NetworkAdapterType *string `yaml:"network_adapter_type,omitempty"`
    Firmware           *string `yaml:"firmware,omitempty"`
}

func (o *VSphereHypervisorDescriptor) String() string {
    bytes, err := yaml.Marshal(o)
    if err != nil {
        return ""
    }
    return string(bytes)
}

func (o *VSphereHypervisorDescriptor) Validate() error {
    return nil
}

func (o *VSphereHypervisorDescriptor) GetGuestOSType() string {
    if o.GuestOSType == nil {
        return ""
    }
    return *o.GuestOSType
}

func (o *VSphereHypervisorDescriptor) GetDiskControllerType() string {
    if o.DiskControllerType == nil {
        return "lsilogic"
    }
    return *o.DiskControllerType
}

func (o *VSphereHypervisorDescriptor) GetNetworkAdapterType() string {
    if o.NetworkAdapterType == nil {
        return "vmxnet3"
    }
    return *o.NetworkAdapterType
}

func (o *VSphereHypervisorDescriptor) GetFirmware() string {
    if o.Firmware == nil {
        return "bios"
    }
    return *o.Firmware
}

func (o *VSphereHypervisorDescriptor) GetName() string {
    return "VMware vSphere"
}

func (o *VSphereHypervisorDescriptor) GetId() string {
    return "vsphere"
}
