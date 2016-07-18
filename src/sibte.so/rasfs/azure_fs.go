package rasfs

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"path"
	"time"

	"github.com/Azure/azure-sdk-for-go/storage"
)

// AzureStorageConfig holds configurations for azure storage account
type AzureStorageConfig struct {
	AccountName string `json:"account_name,omitempty"`
	AccountKey  string `json:"account_key,omitempty"`
	Container   string `json:"container,omitempty"`
	Domain      string `json:"domain,omitempty"`
}

type azureFS struct {
	config *AzureStorageConfig
}

// LoadAzureStorageConfig loads azure storage configuration from given dictionary
func LoadAzureStorageConfig(cfg map[string]string) (*AzureStorageConfig, error) {
	if len(cfg) == 0 {
		return nil, nil
	}

	log.Println("Uploader config...", cfg)

	if cfg["provider"] != "azure" {
		return nil, nil
	}

	jsonBytes, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	cObj := &AzureStorageConfig{}
	if err = json.Unmarshal(jsonBytes, cObj); err != nil {
		return nil, err
	}

	return cObj, nil
}

// NewAzureFS creates instance of AzureFS implementation
func NewAzureFS(cfg *AzureStorageConfig) RasFS {
	return &azureFS{
		config: cfg,
	}
}

// Upload files
func (a *azureFS) Upload(name string, size uint64, reader io.Reader) (string, error) {
	c, err := storage.NewBasicClient(a.config.AccountName, a.config.AccountKey)
	if err != nil {
		return "", err
	}

	uploadPath := a.generateUploadPath(name)
	s := c.GetBlobService()
	err = s.CreateBlockBlobFromReader(a.config.Container, uploadPath, size, reader, nil)
	if err == nil {
		return fmt.Sprintf("http://%s/%s/%s", a.config.Domain, a.config.Container, uploadPath), nil
	}

	return "", err
}

func (a *azureFS) generateUploadPath(name string) string {
	now := time.Now()
	hasher := md5.New()
	io.WriteString(hasher, fmt.Sprintf("%d-%s-%d", rand.Int63(), name, now.Unix()))
	return fmt.Sprintf("%x/%s", hasher.Sum(nil), path.Base(name))
}
