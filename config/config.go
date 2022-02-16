package config

import (
	"path/filepath"
)

type Config struct {
	Excludes    ArrayFlags
	Src         string
	Dest        string
	Rotate      int
	IsRecursive string //if bool, can't judg user input, because bool init to false
	IsDaemon    string
}

//convert config
func mutate(config *Config) error {
	absDest, err := filepath.Abs(config.Dest)
	if err != nil {
		return err
	}
	absSrc, err := filepath.Abs(config.Src)
	if err != nil {
		return err
	}
	config.Dest = absDest
	config.Src = absSrc
	return nil
}

func join(argC *Config, yamlC *Config) (*Config, error) {
	var ar [2]*Config = [2]*Config{yamlC, argC} // overwrite yaml by arg
	joinC := newDefaultConfig()
	for _, c := range ar {
		if c.Dest != "" {
			joinC.Dest = c.Dest
		} else if c.Src != "" {
			joinC.Src = c.Src
		} else if c.Rotate != 0 {
			joinC.Rotate = c.Rotate
		} else if c.IsRecursive != "" {
			joinC.IsRecursive = c.IsRecursive
		} else if len(c.Excludes) != 0 {
			joinC.Excludes = c.Excludes
		} else if c.IsDaemon != "" {
			joinC.IsDaemon = c.IsDaemon
		}
	}

	if err := mutate(joinC); err != nil {
		return nil, err
	}

	return joinC, nil
}

func NewConfig() (*Config, error) {
	yamlC, err := LoadYamlConfig("./config.yml")
	if err != nil {
		return nil, err
	}
	argC := &args
	joinC, err := join(argC, yamlC)
	if err != nil {
		return nil, err
	}
	return joinC, err
}
