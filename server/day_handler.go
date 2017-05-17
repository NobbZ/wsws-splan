package server

import (
	"net/http"

	"log"

	"github.com/julienschmidt/httprouter"

	"github.com/NobbZ/wsws-splan/splan"
)

func daysGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	days := splan.GetAllDays(0)
	log.Println("In daysGetHandler")
	enc := GetEncoderFromRequest(r, w)

	enc.Encode(days)
}

func dayGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("In dayGetHandler")
	w.WriteHeader(http.StatusOK)
}
