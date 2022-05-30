package config

// Configurations exported
type Configurations struct {
	Endpoint         ServerConfigurations
	GIT_ACCESS_TOKEN string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Hostname string
}
