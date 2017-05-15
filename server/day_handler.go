package server

import "net/http"
import "strings"
import "log"

const dayPath = `/day/`

func dayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "OPTIONS":
		dayOptionsHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func dayOptionsHandler(w http.ResponseWriter, r *http.Request) {
	resID := strings.TrimPrefix(r.URL.Path, dayPath)

	log.Fatal("resID:", resID)
}
