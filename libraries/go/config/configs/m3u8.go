package configs

type M3u8Config struct {
	StreamsDirectory string `toml:"streams_directory"`
}

type M3u8 struct {
	Inner M3u8Config `toml:"m3u8"`
}
