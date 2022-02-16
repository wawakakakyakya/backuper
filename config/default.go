package config

func newDefaultConfig() *Config {
	return &Config{Src: "./", Dest: "./", Rotate: 5, IsRecursive: "true", IsDaemon: "false"}
}
