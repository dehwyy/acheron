package configs

type Ports struct {
	Nexus                 uint `toml:"nexus"`
	SrtServer             uint `toml:"srt_server"`
	StreamBroadcasterPort uint `toml:"stream_broadcaster"`
}

type Addr struct {
	Ports Ports `toml:"ports"`
}
