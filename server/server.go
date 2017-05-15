package server

import "net/http"

func Start(addr string) {
	http.HandleFunc(dayPath, dayHandler)
	http.ListenAndServe(addr, nil)
}
