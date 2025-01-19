package config

import (
	"fmt"

	"github.com/dehwyy/mugen/libraries/go/config/configs"
	"github.com/dehwyy/mugen/libraries/go/config/parser"
)

const (
	tomlConfigFilepath = "config/config.toml"
	envConfigFilepath  = ".env"
)

type globalConfig struct {
	addr *configs.Addr
	m3u8 *configs.M3u8
	env  *configs.EnvConfig
}

func (c *globalConfig) Addr() *configs.Addr {
	return c.addr
}

func (c *globalConfig) Env() *configs.EnvConfig {
	return c.env
}

func (c *globalConfig) M3u8() *configs.M3u8Config {
	return &c.m3u8.Inner
}

func (c *globalConfig) String() string {
	return fmt.Sprintf("addr: %+v, m3u8: %+v, env: %+v", c.Addr(), c.M3u8(), c.Env())
}

type ConfigConstructorParams struct {
	EnvFilePath        string `tags:"optional"`
	TomlConfigFilePath string `tags:"optional"`
}

type Opts = ConfigConstructorParams

func New(params ConfigConstructorParams) func() Config {
	tomlFilepath := or(params.TomlConfigFilePath, tomlConfigFilepath)

	return func() Config {
		return &globalConfig{
			addr: parser.ReadAndParse[configs.Addr](tomlFilepath),
			m3u8: parser.ReadAndParse[configs.M3u8](tomlFilepath),
			env:  configs.NewEnvConfig(or(params.EnvFilePath, envConfigFilepath)),
		}
	}
}

func or(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
