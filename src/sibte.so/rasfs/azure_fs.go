package rasfs

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "path"

    "github.com/Azure/azure-sdk-for-go/storage"
)

type azureStorageConfig struct {
    AccountName string `json:"account_name,omitempty"`
    AccountKey  string `json:"account_key,omitempty"`
    Container   string `json:"container,omitempty"`
    Domain      string `json:"domain,omitempty"`
}

type azureFS struct {
    config *azureStorageConfig
}

// NewAzureFS creates instance of AzureFS implementation
func NewAzureFS() RasFS {
    return &azureFS{}
}

func (a *azureFS) Init(cfg map[string]string) error {
    acfg, err := loadAzureStorageConfig(cfg)
    if err != nil {
        return err
    }

    a.config = acfg
    return nil
}

// Upload files
func (a *azureFS) Upload(name string, size uint64, reader io.Reader) (string, error) {
    c, err := storage.NewBasicClient(a.config.AccountName, a.config.AccountKey)
    if err != nil {
        return "", err
    }

    uploadPath := generateUploadPathFromName(name) + path.Base(name)
    s := c.GetBlobService()
    err = s.CreateBlockBlobFromReader(a.config.Container, uploadPath, size, reader, nil)
    if err == nil {
        return fmt.Sprintf("http://%s/%s/%s", a.config.Domain, a.config.Container, uploadPath), nil
    }

    return "", err
}

// LoadAzureStorageConfig loads azure storage configuration from given dictionary
func loadAzureStorageConfig(cfg map[string]string) (*azureStorageConfig, error) {
    if len(cfg) == 0 {
        return nil, nil
    }

    log.Println("Loading azure config...", cfg)

    if cfg["provider"] != "azure" {
        return nil, InvalidConfigurationName
    }

    jsonBytes, err := json.Marshal(cfg)
    if err != nil {
        return nil, err
    }

    cObj := &azureStorageConfig{}
    if err = json.Unmarshal(jsonBytes, cObj); err != nil {
        return nil, err
    }

    return cObj, nil
}
