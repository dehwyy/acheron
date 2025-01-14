package configs

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type EnvConfig struct {
	SentryDsn string `required:"false"   envconfig:"SENTRY_DSN"`
}

func NewEnvConfig(envFilepath string) *EnvConfig {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	_ = godotenv.Load(filepath.Join(wd, envFilepath))
	var config EnvConfig

	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	return &config
}
