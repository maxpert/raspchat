package rasfs

import (
    "encoding/json"
    "io"
    "log"
    "os"
    "path"
    "net/url"
)

type localStorageConfig struct {
    DiskStoragePath string `json:"disk_storage_path,omitempty"`
}

type localFS struct {
    config *localStorageConfig
}

func NewLocalFS() RasFS {
    return &localFS{}
}

func loadLocalStorageConfig(cfg map[string]string) (*localStorageConfig, error) {
    if len(cfg) == 0 {
        return nil, nil
    }

    log.Println("Loading local storage configuration...", cfg)
    if cfg["provider"] != "local" {
        return nil, InvalidConfigurationName
    }

    jsonBytes, err := json.Marshal(cfg)
    if err != nil {
        return nil, err
    }

    cObj := &localStorageConfig{}
    if err = json.Unmarshal(jsonBytes, cObj); err != nil {
        return nil, err
    }

    return cObj, nil
}

func (f *localFS) Init(cfg map[string]string) error {
    tConfig, err := loadLocalStorageConfig(cfg)
    if err != nil {
        return err
    }

    err = os.MkdirAll(tConfig.DiskStoragePath, os.ModePerm)
    if err != nil {
        return err
    }

    f.config = tConfig
    return nil
}

func (f *localFS) Upload(name string, size uint64, reader io.Reader) (string, error) {
    uploadRelPath := generateUploadPathFromName(name)
    uploadAbsPath := path.Join(f.config.DiskStoragePath, uploadRelPath)
    baseName := path.Base(name)
    if err := os.MkdirAll(uploadAbsPath, os.ModePerm); err != nil {
        return "", err
    }

    out, err := os.Create(path.Join(uploadAbsPath, baseName))
    defer (func() {
        out.Close()
    })()

    if err != nil {
        return "", err
    }

    if _, err = io.Copy(out, reader); err != nil {
        return "", err
    }

    if err = out.Sync(); err != nil {
        return "", err
    }

    return uploadRelPath + "/" + url.QueryEscape(baseName), nil
}

func (f *localFS) Download(filePath string) (io.ReadCloser, error) {
    decodedPath, err := url.QueryUnescape(filePath)
    if err != nil {
        return nil, err
    }

    absPath := path.Join(f.config.DiskStoragePath, decodedPath)
    return os.Open(absPath)
}
