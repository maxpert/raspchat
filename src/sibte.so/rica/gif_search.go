package rica

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
)

type atomic_store struct {
	store *bolt.DB
}

var kvStore *atomic_store

func FindRightGif(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := strings.ToLower(r.FormValue("q"))

	if cache_val, ok := kvStore.get(q); ok {
		w.Write([]byte(cache_val))
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
	kvStore.set(q, string(buf.Bytes()))

	io.Copy(w, buf)
}

func (me *atomic_store) get(key string) (string, bool) {
	tx, err := me.store.Begin(false)
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

func (me *atomic_store) set(key, value string) bool {
	tx, err := me.store.Begin(true)
	if err != nil {
		return false
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("rightgif"))
	return bucket.Put([]byte(key), []byte(value)) == nil && tx.Commit() == nil
}

func initGifCache() {
	db, err := bolt.Open(CurrentAppConfig.DBPath+"/gifstore.bolt", 0600, nil)

	if err != nil {
		log.Println("Gif search error", err.Error())
		return
	}

	tx, err := db.Begin(true)
	if err == nil {
		defer tx.Rollback()
		tx.CreateBucket([]byte("rightgif"))
		tx.Commit()
	}

	kvStore = &atomic_store{
		store: db,
	}
}
