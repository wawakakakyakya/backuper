package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type YamlConfig struct {
	Excludes    ArrayFlags `yaml: "excludes"`
	Src         string     `yaml: "src"`
	Dest        string     `yaml: "dest"`
	Rotate      int        `yaml: "rotate"`
	IsRecursive string     `yaml: "is_recursive"`
}

// var c := Config{Exclude: "", Src: "", Dest: "", Rotate: "", IsRecursive: ""}

func LoadYamlConfig(path string) (*Config, error) {
	f, err := ioutil.ReadFile(path)
	var yc *YamlConfig
	// if no config , return default
	if err != nil {
		return &Config{}, nil
	}

	err = yaml.Unmarshal(f, &yc)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("(%%v)  %v\n", c)
	c := &Config{}
	c.Excludes = &yc.Excludes
	c.Dest = yc.Dest
	c.Src = yc.Src
	c.Rotate = yc.Rotate
	c.IsRecursive = yc.IsRecursive

	return c, nil
}
