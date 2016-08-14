package rasweb

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

type directPagesHandler struct {
    pageCache map[string][]byte
}

// NewDirectPagesHandler initializes direct page route handlers
func NewDirectPagesHandler() RouteHandler {
    return &directPagesHandler{make(map[string][]byte)}
}

func (h *directPagesHandler) Register(r *httprouter.Router) error {
    r.GET("/", h.index)
    r.GET("/_clear", h.clearCache)
    return nil
}

func (h *directPagesHandler) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    h.writeFileToResponse(w, r, "static/index.html")
}

func (h *directPagesHandler) writeFileToResponse(w http.ResponseWriter, r *http.Request, path string) {
    if data, ok := h.pageCache[path]; ok {
        w.Header().Add("X-Cache-Hit", "true")
        w.Write(data)
        return
    }

    data, err := ioutil.ReadFile(path)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    h.pageCache[path] = data
    w.Header().Add("X-Cache-Hit", "false")
    w.Write(data)
}

func (h *directPagesHandler) clearCache(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    h.pageCache = make(map[string][]byte)
    fmt.Fprint(w, "")
}
