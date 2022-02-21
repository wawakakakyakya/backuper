package task

import (
	"backuper/config"
	"backuper/logger"
)

func RunSimple(c *config.Config, l logger.LoggerInterface) error {
	l.Info("Run As No Daemon")
	return run(c, l)
}
