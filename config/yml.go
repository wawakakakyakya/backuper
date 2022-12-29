package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfig struct {
	Excludes    []string `yaml:"excludes"`
	Src         string   `yaml:"src"`
	Dest        string   `yaml:"dest"`
	Rotate      int      `yaml:"rotate"`
	IsRecursive bool     `yaml:"is_recursive"`
}

// var c := Config{Exclude: "", Src: "", Dest: "", Rotate: "", IsRecursive: ""}

func LoadYamlConfig(path string) (*Config, error) {

	yc := YamlConfig{}
	err := yml.Load(path, &yc)
	if err != nil {
		return nil, err
	}
	// fmt.Printf("(%%v)  %v\n", c)
	c := &Config{}
	c.Excludes = yc.Excludes
	c.Dest = yc.Dest
	c.Src = yc.Src
	c.Rotate = yc.Rotate
	c.IsRecursive = yc.IsRecursive

	return c, nil
}
