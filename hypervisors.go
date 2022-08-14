package platform

import "fmt"

type HypervisorKind string

const (
	HypervisorKindVMware = "vSphere"
)

func AllHypervisorKinds() []string {
	return []string{
		HypervisorKindVMware,
	}
}

type Hypervisors map[string]*Hypervisor

func (h Hypervisors) VSphere() *Hypervisor {
	// check if HypervisorKindVMware exists in h and return it
	result, ok := h[HypervisorKindVMware]
	if !ok {
		h[HypervisorKindVMware] = defaultHypervisors[HypervisorKindVMware]
		result = h[HypervisorKindVMware]
	}
	return result
}

func (h Hypervisors) validate() error {
	if h == nil {
		return fmt.Errorf("platform: hypervisors is required")
	}

	if len(h) != 1 {
		return fmt.Errorf("platform: only one hypervisor is supported")
	}

	for _, x := range h {
		err := x.validate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (h Hypervisors) String() string {
	return toYAML(h)
}
