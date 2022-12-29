package config

import (
	yml "github.com/wawakakakyakya/configloader/yml"
)

type YamlConfigs struct {
	Cfgs []YamlConfig `yaml:"lists"`
}

type YamlConfig struct {
	Excludes    []string `yaml:"excludes"`
	Src         string   `yaml:"src"`
	Dest        string   `yaml:"dest"`
	Rotate      int      `yaml:"rotate"`
	IsRecursive bool     `yaml:"is_recursive"`
}

// var c := Config{Exclude: "", Src: "", Dest: "", Rotate: "", IsRecursive: ""}

func LoadYamlConfig(path string) (ConfigArray, error) {

	ycArray := YamlConfigs{}
	err := yml.Load(path, &ycArray)
	if err != nil {
		return nil, err
	}
	var cfgArray = ConfigArray{}
	for _, yc := range ycArray.Cfgs {
		c := &Config{}
		c.Excludes = yc.Excludes
		c.Dest = yc.Dest
		c.Src = yc.Src
		c.Rotate = yc.Rotate
		c.IsRecursive = yc.IsRecursive
		cfgArray = append(cfgArray, c)
	}
	return cfgArray, nil
}
