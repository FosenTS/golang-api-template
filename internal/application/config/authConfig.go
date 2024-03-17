package config

import (
	"golang-api-template/pkg/mysync"
	"path"
	"time"
)

const authCfgFilename = "auth.config.yaml"

type AuthConfig struct {
	Salt         string
	SecretJWTKey string

	jwtLiveTimeSeconds     uint `yaml:"jwtLiveTimeSeconds" env-required:"true"`
	refreshLiveTimeSeconds uint `yaml:"refreshLiveTimeSeconds" env-required:"true"`

	JwtLiveTime     time.Duration
	RefreshLiveTime time.Duration
}

var (
	authConfigInst     = &AuthConfig{}
	loadAuthConfigOnce = mysync.NewOnce()
)

func Auth() AuthConfig {
	loadAuthConfigOnce.Do(func() {
		env := Env()
		readConfig(path.Join(env.ConfigAbsPath, authCfgFilename), authConfigInst)

		// TODO: fix token expires time
		authConfigInst.JwtLiveTime = time.Second * time.Duration(authConfigInst.jwtLiveTimeSeconds)
		authConfigInst.JwtLiveTime = 1800 * time.Second
		authConfigInst.RefreshLiveTime = time.Duration(authConfigInst.refreshLiveTimeSeconds) * time.Second
		authConfigInst.RefreshLiveTime = 40323402 * time.Second
	})

	return *authConfigInst
}
