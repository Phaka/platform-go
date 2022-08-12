package platform

// Processors represents the number of processors and cores required for the operating system
type Processors interface {
	// GetCount returns the number of processors required for the operating system
	GetCount() int
	// GetCores returns the number of cores required for the operating system
	GetCores() int
}
