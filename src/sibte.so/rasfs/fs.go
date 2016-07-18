package rasfs

import (
	"io"
)

// RasFS interface defnining all methods related to FS
type RasFS interface {
	Upload(string, uint64, io.Reader) (string, error)
}
