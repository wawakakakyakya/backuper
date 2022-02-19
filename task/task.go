package task

import (
	"backuper/config"
	"backuper/logger"
)

// type Task__ func(mainCtx context.Context, c *config.Config, l logger.LoggerInterface)

type taskFn func(c *config.Config, l logger.LoggerInterface) error

type Task struct {
	fn taskFn
	c  *config.Config
	l  logger.LoggerInterface
}

func (t *Task) Run() error {
	return t.fn(t.c, t.l)
}

func NewTask(fn taskFn, c *config.Config, l logger.LoggerInterface) *Task {
	return &Task{fn, c, l}
}
