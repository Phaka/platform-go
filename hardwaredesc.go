package platform

import "errors"

type HardwareDescriptor struct {
	Memory     *int
	Storage    *int
	Processors *ProcessorsDescriptor
}

func (h *HardwareDescriptor) Validate() error {
	if h.Memory != nil && *h.Memory < 0 {
		return errors.New("invalid memory")
	}
	if h.Storage != nil && *h.Storage < 0 {
		return errors.New("invalid storage")
	}
	if h.Processors != nil {
		return h.Processors.Validate()
	}
	return nil
}

func (h *HardwareDescriptor) GetMemory() int {
	if h.Memory == nil {
		return 2048
	}
	return *h.Memory
}

func (h *HardwareDescriptor) GetStorage() int {
	if h.Storage == nil {
		return 1024 * 10
	}
	if *h.Storage <= 0 {
		return 1024 * 10
	}
	return *h.Storage
}

var one = 1
var defaultProcessors = &ProcessorsDescriptor{
	Count: &one,
	Cores: &one,
}

func (h *HardwareDescriptor) GetProcessors() Processors {
	if h.Processors == nil {
		return defaultProcessors
	}
	return h.Processors
}
