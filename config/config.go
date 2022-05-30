package config

// Configurations exported
type Configurations struct {
	GITHUBAPI_ENDPOINT string
	GITHUBAPI_TOKEN    string
	GitHubCfg          ApiConfigurations
}

// ServerConfigurations exported
type ApiConfigurations struct {
	ApiEndpoint string
	ApiToken    string
}
