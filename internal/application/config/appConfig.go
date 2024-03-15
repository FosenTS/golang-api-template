package config

import (
	"golang-api-template/pkg/mysync"
	"path"
)

const appCfgFilename = "app.config.yaml"

type AppConfig struct {
	LogLevel string `yaml:"logLevel" env-required:"true"`
}

var (
	appConfigInst     = &AppConfig{}
	loadAppConfigOnce = mysync.NewOnce()
)

func App() AppConfig {
	loadAppConfigOnce.Do(func() {
		env := Env()
		readConfig(path.Join(env.ConfigAbsPath, appCfgFilename), appConfigInst)
	})

	return *appConfigInst
}
