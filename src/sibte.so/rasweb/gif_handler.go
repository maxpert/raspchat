package rasweb

import (
    "bytes"
    "io"
    "log"
    "net/http"
    "strings"

    "sibte.so/rasconfig"

    "github.com/julienschmidt/httprouter"
    "github.com/dgraph-io/badger"
    "path"
    "os"
)

type atomicStore struct {
    store *badger.KV
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
    opts := badger.DefaultOptions
    opts.Dir = path.Join(rasconfig.CurrentAppConfig.DBPath, "gifstore", "keys")
    opts.ValueDir = path.Join(rasconfig.CurrentAppConfig.DBPath, "gifstore", "values")

    if err := createPathIfMissing(opts.Dir); err != nil {
        return err
    }

    if err := createPathIfMissing(opts.ValueDir); err != nil {
        return err
    }

    if db, err := badger.NewKV(&opts); err != nil {
        return err
    } else {
        h.kvStore = &atomicStore{
            store: db,
        }
    }

    return nil
}

func (s *atomicStore) get(key string) (string, bool) {
    pair := badger.KVItem{}
    err := s.store.Get([]byte(key), &pair)
    if err != nil || pair.Value() == nil {
        return "", false
    }

    return string(pair.Value()), true
}

func (s *atomicStore) set(key, value string) bool {
    return s.store.Set([]byte(key), []byte(value)) == nil
}

func createPathIfMissing(path string) error {
    if exists, err := pathExists(path); exists == false {
        return os.MkdirAll(path, os.ModePerm)
    } else {
        return err
    }
}

func pathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}