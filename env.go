package echo

type EnvConfig struct {
	GCPProjectID string `envconfig:"GCP_PROJECT_ID"`
	// Following commands should be set as Cloud Build trigger substitutions.
	APIVersion string `envconfig:"_API_VERSION"`
}
