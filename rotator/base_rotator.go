package rotator

import (
	"backuper/config"
	"backuper/logger"
	"fmt"
)

type baseRotator struct {
	rotatorInterface
	config *config.Config
	logger logger.LoggerInterface
}

func (b *baseRotator) remove(files []string, ar actRotator) error {
	if len(files) <= b.config.Rotate {
		b.logger.Info("No rotation target found")
		return nil
	} else if b.config.Rotate == -1 {
		b.logger.Info("Rotation was canceled by config")
		return nil
	}

	for i, f := range files {
		if len(files)-i == b.config.Rotate {
			break
		}
		if err := ar._remove(f); err != nil {
			b.logger.ErrorS(fmt.Sprintf("remove %s was failed", f))
			b.logger.Error(err)
		} else {
			b.logger.Info(fmt.Sprintf("%s was deleted", f))
		}

	}
	return nil
}

func (b *baseRotator) find(ar actRotator) ([]string, error) {
	b.logger.Info("base find called")
	return ar._find()
}

func (b *baseRotator) _run(ar actRotator) error {
	b.logger.Info("start base rotator _run")
	// pdir := filepath.Dir(r.config.Dest)
	files, err := b.find(ar)
	if err != nil {
		b.logger.ErrorS("find rotate file faied")
		return err
	}
	b.logger.Info(fmt.Sprintf("rotation files: %s", files))
	fmt.Println("rotation targets: ", files)
	b.remove(files, ar)

	return nil
}

func newBaseRotator(config *config.Config, logger logger.LoggerInterface) *baseRotator {
	return &baseRotator{config: config, logger: logger}
}

func NewRotator(config *config.Config, logger logger.LoggerInterface) rotatorInterface {
	return newLocalRotator(config, logger)
}
