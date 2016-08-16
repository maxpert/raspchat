package rasweb

import (
    "bytes"
    "io"
    "log"
    "net/http"
    "strings"

    "sibte.so/rasconfig"

    "github.com/julienschmidt/httprouter"
    "github.com/syndtr/goleveldb/leveldb"
)

type atomicStore struct {
    store *leveldb.DB
}

type gifRouteHandler struct {
    kvStore *atomicStore
}

// NewGifHandler creates a route handler for gif finder
func NewGifHandler() RouteHandler {
    return &gifRouteHandler{}
}

func (h *gifRouteHandler) Register(r *httprouter.Router) error {
    if err := h.initGifCache(); err != nil {
        return err
    }

    r.GET("/gif", h.findGifHandler)
    return nil
}

func (h *gifRouteHandler) findGifHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    q := strings.ToLower(r.FormValue("q"))

    if cacheVal, ok := h.kvStore.get(q); ok {
        w.Write([]byte(cacheVal))
        return
    }

    qreader := strings.NewReader("text=" + q)
    resp, err := http.Post("https://rightgif.com/search/web", "application/x-www-form-urlencoded", qreader)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Println("Error", err)
    }

    buf := bytes.NewBuffer(make([]byte, 0))
    io.Copy(buf, resp.Body)
    h.kvStore.set(q, string(buf.Bytes()))

    io.Copy(w, buf)
}

func (h *gifRouteHandler) initGifCache() error {
    db, err := leveldb.OpenFile(rasconfig.CurrentAppConfig.DBPath+"/gifstore.leveldb", nil)

    if err != nil {
        return err
    }

    h.kvStore = &atomicStore{
        store: db,
    }

    return nil
}

func (s *atomicStore) get(key string) (string, bool) {
    ret, err := s.store.Get([]byte(key), nil)
    if err != nil || ret == nil {
        return "", false
    }

    return string(ret), true
}

func (s *atomicStore) set(key, value string) bool {
    return s.store.Put([]byte(key), []byte(value), nil) == nil
}
