package platform

// BootMethod represents the steps to boot an operating system
type BootMethod interface {
	// GetCommands returns the commands to execute to boot the operating system
	GetCommands() string

	// GetFiles returns the files to copy to the boot medium
	GetFiles() map[string]string

	// GetWait returns the amount of time to wait before typing the commands
	GetWait() string

	// GetVariables returns the variables to set when generating the response file
	GetVariables() []Variable
}
