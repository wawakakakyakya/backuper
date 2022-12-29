package archives

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"backuper/config"
	"backuper/logger"
)

type Tar struct {
	buf    *bytes.Buffer
	time   string
	config *config.Config
	logger logger.LoggerInterface
	ArchiveInterface
}

func NewTar(buf *bytes.Buffer, time string, config *config.Config, logger logger.LoggerInterface) *Tar {
	return &Tar{buf: buf, time: time, config: config, logger: logger}
}

func (t *Tar) baseName() string {
	return fmt.Sprintf("%s.%s", filepath.Join(t.config.Dest, filepath.Base(t.config.Src)), t.time)
}

func (t *Tar) Create() error {
	fName := t.baseName() + ".tgz"
	fWriter, err := os.OpenFile(fName, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		t.logger.ErrorS("file open error")
		return err
	}
	defer fWriter.Close()

	if _, err := io.Copy(fWriter, t.buf); err != nil {
		t.logger.ErrorS("write to tarfile error")
		return err
	}
	t.logger.Info(fmt.Sprintf("%s was created", fName))
	return nil
}

func (t *Tar) Add(buf *bytes.Buffer) error {
	zwriter := gzip.NewWriter(buf)
	defer zwriter.Close()

	twriter := tar.NewWriter(zwriter)
	defer twriter.Close()

	t.logger.Info("start Add file")
	t.logger.Info(fmt.Sprintf("src: %s", t.config.Src))
	if err := filepath.Walk(t.config.Src, func(fullPath string, fi os.FileInfo, err error) error {
		if err != nil {
			t.logger.ErrorS("walk file error")
			return err
		}

		// ディレクトリは無視
		if fi.IsDir() {
			return nil
		}
		t.logger.Info(fmt.Sprintf("path: %s", fullPath))
		// ヘッダを書き込み
		relPath, err := filepath.Rel(t.config.Src, fullPath)
		if err != nil {
			t.logger.ErrorS("convert relpath to abspath failed")
			return err
		}
		t.logger.Info(relPath)
		t.logger.Info(fmt.Sprint(filepath.Base(t.config.Src), ".", t.time, "/", relPath))
		if err := twriter.WriteHeader(&tar.Header{
			Name:    fmt.Sprint(filepath.Base(t.config.Src), ".", t.time, "/", relPath),
			Mode:    int64(fi.Mode()),
			ModTime: fi.ModTime(),
			Size:    fi.Size(),
		}); err != nil {
			t.logger.ErrorS("write header error")
			return err
		}

		// ファイルを書き込み
		f, err := os.Open(fullPath)
		if err != nil {
			t.logger.ErrorS(fmt.Sprintf("file(%s) open error", fullPath))
			return err
		}
		defer f.Close()

		if _, err := io.Copy(twriter, f); err != nil {
			t.logger.ErrorS(fmt.Sprintf("(%s) copy file error", fullPath))
			t.logger.Error(err)
			return err
		}
		return nil

	}); err != nil {
		t.logger.Error(err)
		return err
	}
	return nil
}
