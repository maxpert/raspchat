package rica

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/steveyen/gkvlite"
)

type atomic_store struct {
	sync.Mutex
	store *gkvlite.Store
}

var kvStore *atomic_store

func FindRightGif(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := strings.ToLower(r.FormValue("q"))

	if cache_val, ok := kvStore.get(q); ok {
		w.Write([]byte(cache_val))
		return
	}

	log.Println("Searching...", q)
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
	me.Lock()
	defer me.Unlock()

	c := me.store.GetCollection("rightgif")

	ret, err := c.Get([]byte(key))
	return string(ret), (err != nil)
}

func (me *atomic_store) set(key, value string) bool {
	me.Lock()
	defer me.Unlock()

	c := me.store.GetCollection("rightgif")

	err := c.Set([]byte(key), []byte(value))
	return err == nil && me.store.Flush() == nil
}

func InitGifCache(dbPath string) {
	f, _ := os.OpenFile(path.Join(dbPath, "gif.cache"), os.O_RDWR|os.O_CREATE, 0666)
	s, err := gkvlite.NewStore(f)

	if err != nil {
		log.Println("Gif search error", err.Error())
		return
	}

	s.SetCollection("rightgif", nil)
	kvStore = &atomic_store{
		store: s,
	}
}
