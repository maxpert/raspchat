package rasfs

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var InvalidConfigurationName = errors.New("Invalid configuration name")

// RasFS interface defnining all methods related to FS
type RasFS interface {
	Init(map[string]string) error
	Upload(string, uint64, io.Reader) (string, error)
}

type DownloadableRasFS interface {
	Download(string) (io.ReadCloser, error)
}

func generateUploadPathFromName(name string) string {
	now := time.Now()
	hasher := md5.New()
	io.WriteString(hasher, fmt.Sprintf("%d-%s-%d", rand.Int63(), name, now.Unix()))
	return fmt.Sprintf("%d%03d/%x", now.Year(), now.YearDay(), hasher.Sum(nil))
}
