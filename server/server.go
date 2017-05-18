package server

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/NobbZ/wsws-splan/settings"
)

func Start(addr string, wg *sync.WaitGroup) {
	log.Printf("Starting the server")
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	router.HandleOPTIONS = true
	router.HandleMethodNotAllowed = true
	router.GET("/day", daysGetHandler)
	router.GET("/day/:name", dayGetHandler)

	router.GET("/coffee", coffeeGetHandler)
	router.PUT("/coffee", coffeePutHandler)

	if wg != nil {
		wg.Done()
	}

	// http.HandleFunc(dayPath, dayHandler)
	settings.SetBaseURI(addr)

	log.Printf("Listening on %v:%v", getIP(addr), getPort(addr))
	http.ListenAndServe(addr, router)
}

func getIP(addr string) string {
	if addr == "" {
		addr = ":"
	}
	result := strings.Split(addr, ":")
	if result[0] == "" {
		result[0] = "0.0.0.0"
	}
	return result[0]
}

func getPort(addr string) string {
	if addr == "" {
		addr = "http"
	}
	result := strings.Split(addr, ":")
	if result[1] == "" {
		result[1] = "80"
	}
	return result[1]
}
