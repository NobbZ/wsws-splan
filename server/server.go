package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Start(addr string) {
	router := httprouter.New()
	router.HandleOPTIONS = true
	router.GET("/day/:name", dayGetHandler)

	// http.HandleFunc(dayPath, dayHandler)
	http.ListenAndServe(addr, router)
}
