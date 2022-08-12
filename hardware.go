package platform

// Hardware represents the hardware requirements for an operating system
type Hardware interface {
	// GetMemory returns the amount of memory in MB required for the operating system
	GetMemory() int
	// GetStorage returns the amount of storage in GB required for the operating system
	GetStorage() int
	// GetProcessors returns the number of processors and cores required for the operating system
	GetProcessors() Processors
	// Validate returns an error if the hardware specifications are valid
	Validate() error
}
