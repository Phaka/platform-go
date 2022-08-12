package platform

// BootMethods represents the various unattended boot methods available for an operating system
type BootMethods interface {
	// GetHttp returns the steps tp boot the operating system via HTTP
	// This method allows the response file to be stored on an HTTP server and
	// downloaded by the operating system during installation.
	GetHttp() BootMethod
	// GetCdrom returns the steps to boot the operating system via CD-ROM
	// This method allows the response file to be stored on the CD-ROM
	GetCdrom() BootMethod
	// GetFloppy returns the steps to boot the operating system via Floppy
	// This method allows the response file to be stored on the Floppy
	GetFloppy() BootMethod
}
