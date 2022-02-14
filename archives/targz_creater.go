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
	today  string
	config *config.Config
	logger logger.LoggerInterface
	ArchiveInterface
}

func NewTar(buf *bytes.Buffer, today string, config *config.Config, logger logger.LoggerInterface) *Tar {
	return &Tar{buf: buf, today: today, config: config, logger: logger}
}

func (t *Tar) Create() error {
	fName := fmt.Sprintf("%s.%s.tgz", filepath.Join(t.config.Dest, filepath.Base(t.config.Src)), t.today)
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
	if err := filepath.Walk(t.config.Src, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			t.logger.ErrorS("walk file error")
			return err
		}

		// ディレクトリは無視
		if fi.IsDir() {
			return nil
		}

		// ヘッダを書き込み
		if err := twriter.WriteHeader(&tar.Header{
			Name:    path,
			Mode:    int64(fi.Mode()),
			ModTime: fi.ModTime(),
			Size:    fi.Size(),
		}); err != nil {
			t.logger.ErrorS("write header error")
			return err
		}

		// ファイルを書き込み
		f, err := os.Open(path)
		if err != nil {
			t.logger.ErrorS("file open error")
			return err
		}
		defer f.Close()

		if _, err := io.Copy(twriter, f); err != nil {
			t.logger.ErrorS("copy file error")
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
