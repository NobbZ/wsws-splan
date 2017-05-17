package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/NobbZ/wsws-splan/splan"
)

func daysGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	days := splan.GetAllDays(0)
	enc := GetEncoderFromRequest(r, w)

	enc.Encode(days)
}

func dayGetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
