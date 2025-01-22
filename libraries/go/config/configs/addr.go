package configs

type Ports struct {
	Nexus                 uint `toml:"nexus"`
	SrtServer             uint `toml:"srt_server"`
	StreamBroadcasterPort uint `toml:"stream_broadcaster"`
	StreamWhip            uint `toml:"stream_whip"`
}

type Addr struct {
	Ports Ports `toml:"ports"`
}
