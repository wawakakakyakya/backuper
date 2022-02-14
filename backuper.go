package main

import (
	"backuper/archives"
	"backuper/config"
	"backuper/logger"
	"backuper/rotator"
	"bytes"
	"os"
	"time"
)

var (
	OK = 0
	NG = 1
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

func exitWithError(logger logger.LoggerInterface, err error) {
	logger.Error(err)
	os.Exit(NG)
}

func main() {
	// flag.Parse()
	logger := logger.NewLogger()
	logger.Info("start main")
	config := config.NewConfig()
	rotator := rotator.NewRotator(config, logger)
	// fmt.Printf("exclude: %v\n", libs.Args.Excludes)
	// fmt.Printf("src: %v\n", libs.Args.Src)
	// fmt.Printf("dest: %v\n", libs.Args.Dest)

	today, err := getDate()
	if err != nil {
		logger.ErrorS("get date error")
		exitWithError(logger, err)
	}
	_, sErr := os.Stat(config.Src)
	if os.IsNotExist(sErr) {
		logger.ErrorS("src does not exist")
		exitWithError(logger, sErr)
	}

	if err = makeDestDir(config); err != nil {
		exitWithError(logger, err)
	}
	//fmt.Printf("src: %p", &libs.Args.Src)
	buf := new(bytes.Buffer)
	tar := archives.NewTar(buf, today, config, logger)
	if err := tar.Add(buf); err != nil {
		exitWithError(logger, err)
	}

	if err := tar.Create(); err != nil {
		exitWithError(logger, err)
	}
	// act.Run -> base._run(act) -> base.find(act) -> act._find
	// act.Run -> base._runのときにactを渡すことでinterfaceを満たしている
	rotator.Run()
	os.Exit(OK)
}
