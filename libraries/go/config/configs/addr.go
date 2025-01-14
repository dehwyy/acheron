package configs

import (
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type Ports struct {
	SrtServer             uint `toml:"srt_server"`
	StreamBroadcasterPort uint `toml:"stream_broadcaster"`
}

type Addr struct {
	Ports Ports
}

func NewAddrConfig(tomlFilepath string) *Addr {
	file, err := os.Open(tomlFilepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Addr

	if err := toml.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	return &config
}
