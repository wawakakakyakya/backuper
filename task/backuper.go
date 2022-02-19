package task

import (
	"backuper/archives"
	"backuper/config"
	"backuper/logger"
	"backuper/rotator"
	"bytes"
	"fmt"
	"os"
	"time"
)

func getDate() (string, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	nowJST := time.Now().In(jst)
	return nowJST.Format("20060102"), nil
}

func makeDestDir(config *config.Config) error {
	fInfo, err := os.Stat(string(config.Dest))
	if err != nil {
		return err
	}
	if !fInfo.IsDir() {
		os.MkdirAll(string(config.Dest), os.ModePerm)
	}
	return nil
}

func run(c *config.Config, l logger.LoggerInterface) error {
	// flag.Parse()
	logger := logger.NewLogger()
	logger.Info("start main")
	fmt.Printf("(%%v)  %v\n", c)
	rotator := rotator.NewRotator(c, logger)
	// fmt.Printf("exclude: %v\n", libs.Args.Excludes)
	// fmt.Printf("src: %v\n", libs.Args.Src)
	// fmt.Printf("dest: %v\n", libs.Args.Dest)

	today, err := getDate()
	if err != nil {
		logger.ErrorS("get date error")
		return err
	}
	_, sErr := os.Stat(c.Src)
	if os.IsNotExist(sErr) {
		logger.ErrorS("src does not exist")
		return err
	}

	if err = makeDestDir(c); err != nil {
		return err
	}
	//fmt.Printf("src: %p", &libs.Args.Src)
	buf := new(bytes.Buffer)
	tar := archives.NewTar(buf, today, c, logger)
	if err := tar.Add(buf); err != nil {
		return err
	}

	if err := tar.Create(); err != nil {
		return err
	}
	// act.Run -> base._run(act) -> base.find(act) -> act._find
	// act.Run -> base._runのときにactを渡すことでinterfaceを満たしている
	rotator.Run()
	return nil
}
