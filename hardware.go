package platform

type Hardware interface {
	GetMemory() int
	GetStorage() int
	GetProcessors() Processors
	Validate() error
}
