package handler

import (
	"backuper/config"
	"backuper/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SigTermHandler(l logger.LoggerInterface, cancel context.CancelFunc) {
	//use buffered channel
	//https://budougumi0617.github.io/2020/09/06/why_signal_notify_want_buffered_channel/
	endChannel := make(chan os.Signal, 1)
	defer close(endChannel)
	signal.Notify(endChannel, syscall.SIGINT, syscall.SIGTERM)

	//stop proccess by stop/ctrl-c
	<-endChannel // block func
	l.Info("signal INT/TERM received., stop context")
	cancel() //send true to ctx.Done channel
}

func SigHupHandler(mainCtx context.Context, l logger.LoggerInterface) {
	reloadChannel := make(chan os.Signal, 1)
	defer close(reloadChannel)
	signal.Notify(reloadChannel, syscall.SIGHUP)

	//Don't stop if error
	for {
		select {
		case <-reloadChannel:
			if err := config.ReloadConfig(); err != nil {
				l.ErrorS("Reload config failed")
				l.Error(err)
			} else {
				l.Info("Reload config succeeded")
				l.Info(fmt.Sprintf("%v", config.GlobalConfig))
			}
		case <-mainCtx.Done():
			return
		}
	}
}
