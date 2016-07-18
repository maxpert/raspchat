package rasweb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"sibte.so/rasconfig"
	"sibte.so/rasfs"
)

type fileUploadHandler struct {
	fsUploader rasfs.RasFS
}

// NewFileUploadHandler handles file upload requests
func NewFileUploadHandler() RouteHandler {
	return &fileUploadHandler{}
}

func (p *fileUploadHandler) Register(r *httprouter.Router) error {
	cfg, err := rasfs.LoadAzureStorageConfig(rasconfig.CurrentAppConfig.UploaderConfig)
	if err != nil {
		return err
	}

	// Incase no configuration is found don't hook any endpoints
	if cfg == nil {
		return nil
	}

	log.Println("Hooking files routes...")
	p.fsUploader = rasfs.NewAzureFS(cfg)
	r.POST("/file", p.upload)
	r.PUT("/file", p.upload)
	return nil
}

func (p *fileUploadHandler) upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Defer printing error
	var err error
	defer func() {
		if err == nil {
			return
		}
		w.WriteHeader(500)
		fmt.Fprintf(w, "Unable to process file upload error: %s", err.Error())
	}()

	if err = r.ParseMultipartForm(32 << 10); err != nil {
		return
	}

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		return
	}

	defer uploadedFile.Close()
	fileSize, err := uploadedFile.Seek(0, os.SEEK_END)
	if err != nil {
		return
	}

	_, err = uploadedFile.Seek(0, os.SEEK_SET)
	if err != nil {
		return
	}

	url, err := p.fsUploader.Upload(handler.Filename, uint64(fileSize), uploadedFile)
	if err != nil {
		return
	}

	response, err := json.Marshal(struct {
		URL string `json:"url"`
	}{
		URL: url,
	})

	if err != nil {
		return
	}

	w.Write(response)
}
