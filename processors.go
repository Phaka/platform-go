package platform

type Processors interface {
	GetCount() int
	GetCores() int
}
