package config

import (
	"flag"
	"fmt"
	"path/filepath"
)

type AbsPath string

func (ap *AbsPath) String() string {
	return ""
}

func (ap *AbsPath) Set(v string) error {
	// p, err := filepath.Abs(string(*ap))
	p, err := filepath.Abs(v)
	fmt.Println("abs path: ", p)
	if err != nil {
		fmt.Println("get file Abs path error")
		return err
	}
	*ap = AbsPath(p)
	return nil
}

// func (ap *AbsPath) IsSubDir(v string) error {
// 	filepath.Rel()
// }

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
	*af = append(*af, v)
	return nil
}

var args Config

func init() {
	fmt.Println("start libs.args.init")
	flag.IntVar(&args.MaxLength, "length", 24, "string lentgh")
	flag.Var((*AbsPath)(&args.Src), "src", "/root/default")
	flag.Var((*AbsPath)(&args.Dest), "dest", "/root/default")
	flag.Var((*ArrayFlags)(&args.Excludes), "exclude", "Some description for this param.")
	flag.IntVar(&args.Rotate, "rotate", 5, "rotate backup file")
	flag.BoolVar(&args.IsRecursive, "recursive", true, "backup subdir")
	flag.Parse()

	fmt.Println("libs.args.init")
	fmt.Println("args:")
	fmt.Println("src: ", args.Src)
	fmt.Println("destg: ", args.Dest)
	fmt.Println("exclude: ", args.Excludes)
}
