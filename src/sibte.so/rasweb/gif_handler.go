package rasweb

import (
    "bytes"
    "io"
    "log"
    "net/http"
    "strings"

    "github.com/boltdb/bolt"
    "github.com/julienschmidt/httprouter"

    "sibte.so/rasconfig"
    "sibte.so/rascore/utils"
    "path"
)

var _DefaultBucket []byte = []byte("gif")

type atomicStore struct {
    store *bolt.DB
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

    qReader := strings.NewReader("text=" + q)
    resp, err := http.Post("https://rightgif.com/search/web", "application/x-www-form-urlencoded", qReader)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Println("Error", err)
    }

    log.Println("Saving back gif...")
    buf := bytes.NewBuffer(make([]byte, 0))
    io.Copy(buf, resp.Body)
    h.kvStore.set(q, string(buf.Bytes()))

    log.Println("Streaming back gif...")
    io.Copy(w, buf)
}

func (h *gifRouteHandler) initGifCache() error {
    err := rasutils.CreatePathIfMissing(rasconfig.CurrentAppConfig.DBPath)
    if err != nil {
        return err
    }

    dbPath := path.Join(rasconfig.CurrentAppConfig.DBPath, "gif.db")
    if db, err := bolt.Open(dbPath, 0660, nil); err != nil {
        return err
    } else if db != nil {
        h.kvStore = &atomicStore{
            store: db,
        }
    }

    return nil
}

func (s *atomicStore) get(key string) (string, bool) {
    tnx, err := s.store.Begin(false)
    if err != nil {
        return "", false
    }
    defer tnx.Rollback()

    buck := tnx.Bucket(_DefaultBucket)
    if buck == nil {
        return "", false
    }

    val := buck.Get([]byte(key))
    if val == nil {
        return "", false
    }

    return string(append([]byte(nil), val...)), true
}

func (s *atomicStore) set(key, value string) bool {
    tnx, err := s.store.Begin(true)
    if err != nil {
        return false
    }

    defer tnx.Commit()
    buck, err := tnx.CreateBucketIfNotExists(_DefaultBucket)
    if err != nil {
        return false
    }

    err = buck.Put([]byte(key), []byte(value))
    if err != nil {
        return false
    }

    return true
}
