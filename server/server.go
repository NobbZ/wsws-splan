package server

import (
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/NobbZ/wsws-splan/settings"
)

func Start(addr string, wg *sync.WaitGroup) {
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.HandleOPTIONS = true
	router.HandleMethodNotAllowed = true
	router.GET("/day", daysGetHandler)
	router.GET("/day/:name", dayGetHandler)

	if wg != nil {
		wg.Done()
	}

	// http.HandleFunc(dayPath, dayHandler)
	settings.SetBaseURI(addr)
	http.ListenAndServe(addr, router)
}
