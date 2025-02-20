package config

import (
	"fmt"

	"github.com/dehwyy/acheron/libraries/go/config/configs"
)

type Config interface {
	fmt.Stringer

	Addr() *configs.Addr
	M3u8() *configs.M3u8Config
	Env() *configs.EnvConfig
}
