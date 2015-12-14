package rica

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func FindRightGif(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.FormValue("q")
	qreader := strings.NewReader("text=" + q)
	resp, err := http.Post("https://rightgif.com/search", "application/x-www-form-urlencoded", qreader)
	if err != nil {
		fmt.Fprintf(w, "null")
	}

	io.Copy(w, resp.Body)
}
