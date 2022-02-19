package runner

import (
	"backuper/config"
	"backuper/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var endChannel chan os.Signal
var reloadChannel chan os.Signal

//シグナルを受け取ったら、キャンセルさせる

func signalHandler(cancel context.CancelFunc, l logger.LoggerInterface) {

	endChannel = make(chan os.Signal)
	reloadChannel = make(chan os.Signal)
	// signal.Notifyを使ってシグナルを待つ
	signal.Notify(endChannel, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(reloadChannel, syscall.SIGHUP)

	go func() {
		<-endChannel
		fmt.Println("signal INT/TERM received., stop context")
		cancel()
	}()
	//reload config by sighup
	go func() {
		for {
			<-reloadChannel
			fmt.Println("signal received., finish process")
			if err := config.ReloadConfig(); err != nil {
				l.ErrorS("Reload config failed")
				l.Error(err)
			}
		}
	}()
}
