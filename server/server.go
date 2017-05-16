package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Start(addr string) {
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.HandleOPTIONS = true
	router.HandleMethodNotAllowed = true
	router.GET("/day", daysGetHandler)
	router.GET("/day/:name", dayGetHandler)

	// http.HandleFunc(dayPath, dayHandler)
	http.ListenAndServe(addr, router)
}
