package config

import "github.com/dehwyy/mugen/libraries/go/config/configs"

const (
	tomlConfigFilepath = "config/config.toml"
	envConfigFilepath  = ".env"
)

type Config interface {
	Addr() *configs.Addr
	Env() *configs.EnvConfig
}

type GlobalConfig struct {
	addr *configs.Addr
	env  *configs.EnvConfig
}

func (c *GlobalConfig) Addr() *configs.Addr {
	return c.addr
}

func (c *GlobalConfig) Env() *configs.EnvConfig {
	return c.env
}

type Opts struct {
	EnvFilePath        string `tags:"optional"`
	TomlConfigFilePath string `tags:"optional"`
}

func New(opts Opts) func() Config {

	return func() Config {
		return &GlobalConfig{
			addr: configs.NewAddrConfig(or(opts.TomlConfigFilePath, tomlConfigFilepath)),
			env:  configs.NewEnvConfig(or(opts.EnvFilePath, envConfigFilepath)),
		}
	}
}

func or(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
