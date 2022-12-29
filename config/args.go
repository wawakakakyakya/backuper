package config

import (
	"flag"
	"fmt"
)

var configPath string

func init() {
	fmt.Println("start libs.args.init")
	flag.StringVar(&configPath, "config", "config.yml", "config path, default ./config.yml")
	flag.Parse()
}
