package server

import (
	"io"
	"net/http"

	"encoding/json"
	"encoding/xml"

	"bitbucket.org/ww/goautoneg"
)

var alternatives = []string{
	"application/json",
	"application/xml",
}

type Encoder interface {
	Encode(v interface{}) error
}

func GetEncoderFromRequest(r *http.Request, w io.Writer) Encoder {
	accept, ok := r.Header["Accept"]
	if !ok {
		accept = []string{"application/json"}
	}

	switch goautoneg.Negotiate(accept[0], alternatives) {
	case "application/json":
		return json.NewEncoder(w)
	case "application/xml":
		return xml.NewEncoder(w)
	}
	return nil
}
