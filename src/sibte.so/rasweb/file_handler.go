package rasweb

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/julienschmidt/httprouter"
    "sibte.so/rasconfig"
    "sibte.so/rasfs"
)

// MaxFileSizeLimit is 64 MB
const MaxFileSizeLimit = 64 << 20

type fileUploadHandler struct {
    fsUploader   rasfs.RasFS
    fsDownloader rasfs.DownloadableRasFS
}

// NewFileUploadHandler handles file upload requests
func NewFileUploadHandler() RouteHandler {
    return &fileUploadHandler{}
}

func (p *fileUploadHandler) Register(r *httprouter.Router) error {
    configs := []rasfs.RasFS{
        rasfs.NewLocalFS(),
    }

    for _, fs := range configs {
        err := fs.Init(rasconfig.CurrentAppConfig.UploaderConfig)
        if err == nil {
            p.fsUploader = fs
            break
        }

        log.Println("Error fs.Init", err)
    }

    if p.fsUploader == nil {
        return nil
    }

    log.Println("Hooking files routes...")
    r.POST("/file", p.upload)
    r.PUT("/file", p.upload)

    var ok bool
    if p.fsDownloader, ok = p.fsUploader.(rasfs.DownloadableRasFS); ok {
        log.Println("Downloadable file upload handler detected...", p.fsDownloader)
        r.GET("/file/*downloadId", p.download)
    }

    return nil
}

func (p *fileUploadHandler) download(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    if p.fsDownloader == nil {
        w.WriteHeader(422)
        fmt.Fprintf(w, "Invalid uploader")
        return
    }

    reader, err := p.fsDownloader.Download(params.ByName("downloadId"))
    if err != nil {
        w.WriteHeader(404)
        fmt.Fprintf(w, "Unable to process request %v", err)
        return
    }
    defer (func() {
        reader.Close()
    })()

    io.Copy(w, reader)
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
    defer uploadedFile.Close()
    if err != nil {
        return
    }

    defer uploadedFile.Close()
    fileSize, err := uploadedFile.Seek(0, os.SEEK_END)
    if err != nil {
        return
    }

    if fileSize > MaxFileSizeLimit {
        err = errors.New("File size too long")
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

    if p.fsDownloader != nil {
        log.Println("Appending /file/ to", url)
        url = "/file/" + url
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
