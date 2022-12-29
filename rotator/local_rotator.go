package rotator

import (
	"backuper/config"
	"backuper/logger"
	"fmt"
	"os"
	"path/filepath"
)

//ローカルファイルのローテーター
//ロガーは親で持つので持たせない
// Todo: Remove Logger
type localRotator struct {
	*baseRotator
	config *config.Config
	logger logger.LoggerInterface
}

func (l *localRotator) Run() error {
	l.logger.Info("act runner called")
	l._run(l) // input interface
	return nil
}

func (l *localRotator) _remove(fpath string) error {
	l.logger.Info("call base find")
	err := os.Remove(fpath)
	// l.logger.Info(fmt.Sprintf("%s was deleted", fpath))
	return err
}

func (l *localRotator) _find() ([]string, error) {
	l.logger.Info("act rotator _find called")
	pdir := l.config.Dest
	p := filepath.Join(pdir, filepath.Base(l.config.Src)+"*.tgz")
	l.logger.Info(fmt.Sprintf("fullpath: %s", p))
	return filepath.Glob(p)
}

func newLocalRotator(config *config.Config, logger logger.LoggerInterface) rotatorInterface {
	br := newBaseRotator(config, logger)
	return &localRotator{baseRotator: br, config: config, logger: logger}
}
