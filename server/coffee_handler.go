package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func coffeeGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusTeapot)
}

func coffeePutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusTeapot)
}
