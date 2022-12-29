package main

import (
	"backuper/archives"
	"backuper/config"
	"backuper/logger"
	"backuper/rotator"
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	OK = 0
	NG = 1
	wg sync.WaitGroup
)

func getDate() (string, error) {
	// Todo: get localtime
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	nowJST := time.Now().In(jst)
	return nowJST.Format("20060102-030405"), nil
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

func do(l logger.LoggerInterface, c *config.Config, errCh chan<- error) {
	defer wg.Done()
	time, err := getDate()
	if err != nil {
		l.ErrorS("get date error")
		errCh <- err
		return
	}
	_, sErr := os.Stat(c.Src)
	if os.IsNotExist(sErr) {
		l.ErrorS("src does not exist")
		errCh <- err
		return
	}

	if err = makeDestDir(c); err != nil {
		errCh <- err
		return
	}
	//fmt.Printf("src: %p", &libs.Args.Src)
	buf := new(bytes.Buffer)
	tar := archives.NewTar(buf, time, c, l)
	if err := tar.Add(buf); err != nil {
		l.ErrorS("tar.Add failed")
		errCh <- err
		return
	}

	if err := tar.Create(); err != nil {
		l.ErrorS("tar.Create failed")
		errCh <- err
		return
	}
	rotator := rotator.NewRotator(c, l)
	rotator.Run()
	l.Info("wg decrement")

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
	errCh := make(chan error, len(cfgAr))
	defer close(errCh)
	for _, config := range cfgAr {
		wg.Add(1)
		logger.Info("wg add")
		go do(logger, config, errCh)
	}

	wg.Wait()
	select {
	case err, closed := <-errCh:
		if !closed {
			fmt.Printf("Value %s was read.\n", err.Error())
		} else {
			fmt.Println("Channel closed!")
		}
	default:
		fmt.Println("No value ready, moving on.")
	}

	logger.Info("task was ended")
	os.Exit(OK)
}
