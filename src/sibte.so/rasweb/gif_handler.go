package rasweb

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"sibte.so/rasconfig"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
)

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

func (h gifRouteHandler) initGifCache() error {
	db, err := bolt.Open(rasconfig.CurrentAppConfig.DBPath+"/gifstore.bolt", 0600, nil)

	if err != nil {
		return err
	}

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	tx.CreateBucket([]byte("rightgif"))
	tx.Commit()
	h.kvStore = &atomicStore{
		store: db,
	}

	return nil
}

func (s *atomicStore) get(key string) (string, bool) {
	tx, err := s.store.Begin(false)
	if err != nil {
		return "", false
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("rightgif"))

	ret := bucket.Get([]byte(key))
	if ret == nil {
		return "", false
	}

	return string(ret), true
}

func (s *atomicStore) set(key, value string) bool {
	tx, err := s.store.Begin(true)
	if err != nil {
		return false
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("rightgif"))
	return bucket.Put([]byte(key), []byte(value)) == nil && tx.Commit() == nil
}
