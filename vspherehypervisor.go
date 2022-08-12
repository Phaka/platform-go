package platform

type VSphereHypervisor interface {
    GetGuestOSType() string
    GetDiskControllerType() string
    GetNetworkAdapterType() string
    GetFirmware() string
    GetName() string
    GetId() string
}
