package scheduler

import (
	"backuper/logger"
	"backuper/task"
	"context"
	"time"
)

type JobInterface interface {
	AddTask(t *task.Task)
	Run() error
}

type Job struct {
	task *task.Task
	l    logger.LoggerInterface
	ctx  context.Context
	tick *time.Time
	// args interface{}
}

func (j *Job) SetTick(tick *time.Time) {
	j.tick = tick
}

//start task with specified time
func (j *Job) Run() {
	j.l.Info("Run As Daemon")
	timer := time.NewTimer(time.Until(*j.tick))
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			j.l.Info("run started")
			if err := j.task.Run(); err != nil {
				j.l.Error(err)
			}
			j.l.Info("run finished")
		case <-j.ctx.Done(): // called cancel by signal
			j.l.Info("RunAsDaemon ctx.Done called!")
			wg.Done()
			return
			// default:
			// 	l.Info("waiting 1sec...")
			// 	time.Sleep(1 * time.Millisecond)
		}
	}
}

func NewJob(t *task.Task, l logger.LoggerInterface, ctx context.Context) *Job {
	return &Job{task: t, l: l, ctx: ctx}
}
