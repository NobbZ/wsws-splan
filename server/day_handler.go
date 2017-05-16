package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func dayGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
