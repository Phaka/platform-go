package platform

import "fmt"

type BootMethodKind string

const (
	BootMethodKindHttp   = "http"
	BootMethodKindCdrom  = "cdrom"
	BootMethodKindFloppy = "floppy"
)

func BootMethodKinds() []string {
	return []string{
		BootMethodKindHttp,
		BootMethodKindCdrom,
		BootMethodKindFloppy,
	}
}

type BootMethods map[string]*BootMethod

func (b BootMethods) Default() *BootMethod {
	return b.Http()
}

func (b BootMethods) Http() *BootMethod {
	return b[BootMethodKindHttp]
}

func (b BootMethods) Cd() *BootMethod {
	return b[BootMethodKindHttp]
}

func (b BootMethods) Floppy() *BootMethod {
	return b[BootMethodKindHttp]
}

func (b BootMethods) String() string {
	return toYAML(b)
}

func (b BootMethods) validate() error {
	if b == nil {
		return fmt.Errorf("platform: boot methods is required")
	}

	if len(b) == 0 {
		return fmt.Errorf("platform: at least one boot method is required")
	}

	for _, x := range b {
		err := x.validate()
		if err != nil {
			return err
		}
	}
	return nil

}
