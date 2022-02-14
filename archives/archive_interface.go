package archives

import (
	"bytes"
)

type ArchiveInterface interface {
	Create() error
	Add(buf *bytes.Buffer) error
}
