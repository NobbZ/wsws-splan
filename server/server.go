package server

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var BasePath = ""

func Start(addr string) {
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.HandleOPTIONS = true
	router.HandleMethodNotAllowed = true
	router.GET("/day", daysGetHandler)
	router.GET("/day/:name", dayGetHandler)

	// http.HandleFunc(dayPath, dayHandler)
	setBaseURI(addr)
	http.ListenAndServe(addr, router)
}

func setBaseURI(addr string) {
	parts := strings.SplitN(addr, ":", 2)

	if parts[1] == "80" || parts[1] == "http" {
		BasePath = "http://localhost/"
	} else {
		BasePath = "http://localhost:" + parts[1] + "/"
	}
}
