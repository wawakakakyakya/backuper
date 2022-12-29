package main

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

var (
	OK = 0
	NG = 1
)

func getDate() (string, error) {
	// Todo: get localtime
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	nowJST := time.Now().In(jst)
	return nowJST.Format("2006-01-02_03-04-05"), nil
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
	logger := logger.NewLogger()
	logger.Info("start main")
	cfgAr, err := config.NewConfig()
	fmt.Printf("(%%v)  %v\n", cfgAr)
	if err != nil {
		logger.ErrorS("get config failed")
		exitWithError(logger, err)
	}
	for _, config := range cfgAr {
		rotator := rotator.NewRotator(config, logger)

		time, err := getDate()
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
		tar := archives.NewTar(buf, time, config, logger)
		if err := tar.Add(buf); err != nil {
			exitWithError(logger, err)
		}

		if err := tar.Create(); err != nil {
			exitWithError(logger, err)
		}
		rotator.Run()
	}
	os.Exit(OK)
}
