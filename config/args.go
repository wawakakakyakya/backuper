package config

import (
	"flag"
	"fmt"
	"path/filepath"
)

type ArrayFlags []string

func (i *ArrayFlags) String() string {
	//return fmt.Sprint("%v", *i)
	return ""
}

func (af *ArrayFlags) Set(v string) error {
	fmt.Println("exclude in args.go: ", v)
	p, err := filepath.Abs(string(v))
	fmt.Println("abs path in args.go: ", p)
	if err != nil {
		fmt.Println("get file Abs path error")
		return err
	}
	if v != "" {
		*af = append(*af, v)
	}

	return nil
}

var args Config

func init() {
	fmt.Println("start libs.args.init")
	flag.StringVar(&args.Src, "src", "", "default ./")
	flag.StringVar(&args.Dest, "dest", "", "default ./")
	// flag.Var((*ArrayFlags)(args.Excludes), "exclude", "Some description for this param.")
	flag.Var(&args.Excludes, "exclude", "Some description for this param.")
	flag.IntVar(&args.Rotate, "rotate", 0, "rotate backup file (default 5)")
	flag.StringVar(&args.IsRecursive, "recursive", "", "true or false (default true)")
	flag.StringVar(&args.IsDaemon, "daemon", "", "true or false (default false)")
	flag.Parse()

	fmt.Println("libs.args.init")
	fmt.Println("args:")
	fmt.Println("src: ", args.Src)
	fmt.Println("destg: ", args.Dest)
	fmt.Println("exclude: ", args.Excludes)
	fmt.Println("is_daemon: ", args.IsDaemon)
}
