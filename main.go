package main

import (
	"backuper/config"
	"backuper/logger"
	"backuper/scheduler"
	"backuper/task"
	"context"
	"os"
	"sync"
	"time"
)

var (
	OK    int = 0
	NG    int = 1
	eCode int
)

func exitWithError(logger logger.LoggerInterface, err error) {
	logger.Error(err)
	os.Exit(NG)
}

func main() {
	l := logger.NewLogger()
	c, err := config.NewConfig()
	if err != nil {
		l.ErrorS("get config failed")
		exitWithError(l, err)
	}

	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	s := scheduler.NewScheduler(l)
	now := time.Now()
	for i, config := range []*config.Config{} {
		duration := i + 1
		t := task.NewTask(task.RunSimple, config, l)
		job := scheduler.NewJob(t, l, mainCtx)
		if err := s.AddJob(job, now.Add(time.Duration(duration)*time.Second).Format(time.RFC3339)); err != nil {
			continue
		}
		wg.Add(1) //goroutineの中ではAddしない（goroutineが実行されるタイミングはバラバラだから）
	}
	scheduler.Run()

	// var task func(mainCtx context.Context, c *config.Config, l logger.LoggerInterface) error
	// var runTask task.Task
	// if c.IsDaemon == "true" {
	// 	runTask = task.RunAsDaemon
	// } else {
	// 	runTask = task.RunAsDaemon
	// }

	// go handler.SigTermHandler(l, cancel)
	// go handler.SigHupHandler(mainCtx, l)
	// go func() {
	// 	if err = runTask(mainCtx, c, l); err != nil {
	// 		l.Error(err)
	// 		eCode = NG
	// 	} else {
	// 		eCode = OK
	// 	}
	// 	wg.Done() // wg -1
	// }()

	wg.Wait()

	os.Exit(eCode)
}
