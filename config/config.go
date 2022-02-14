package config

type Config struct {
	Excludes    ArrayFlags
	MaxLength   int
	Src         string
	Dest        string
	Rotate      int
	IsRecursive bool
}

func join() string {
	return ""
}

func NewConfig() *Config {
	return &args
}
