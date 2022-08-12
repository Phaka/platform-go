package platform

// Variable represents the variables to set when generating the response file
type Variable interface {
	// GetName returns the name of the variable
	GetName() string
	// GetDefaultValue returns the default value of the variable
	GetDefaultValue() *string
	// GetType returns the type of the variable
	GetType() string
}
