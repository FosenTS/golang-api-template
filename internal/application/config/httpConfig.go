package config

import (
	"golang-api-template/pkg/mysync"
	"path"
)

const httpConfigFilename = "http.config.yaml"

type HTTPConfig struct {
	Host                    string
	Port                    string
	UseCache                bool `yaml:"useCache" env-required:"true"`
	MaxConcurrentConnection uint `yaml:"maxConcurrentConnection" env-required:"true`
}

var (
	httpConfigInst     = &HTTPConfig{}
	loadHTTPConfigOnce = mysync.NewOnce()
)

func HTTP() HTTPConfig {
	loadHTTPConfigOnce.Do(func() {
		env := Env()

		httpConfigInst.Host = env.IpAddress
		httpConfigInst.Port = env.ApiPort
		readConfig(path.Join(env.ConfigAbsPath, httpConfigFilename), httpConfigInst)
	})

	return *httpConfigInst
}
