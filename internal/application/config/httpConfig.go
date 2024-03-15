package config

const httpConfigFilename = "http.config.yaml"

type HTTPConfig struct {
	Host     string
	Port     string
	UseCache bool `yaml:"useCache" env-required:"true"`
}
