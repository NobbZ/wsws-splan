package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestDayOptions(t *testing.T) {
	go Start("localhost:9191")

	req, _ := http.NewRequest("GET", "http://localhost:9191/day", nil)
	c := new(http.Client)
	req.Header["Accept"] = []string{"application/xml"}
	rep, _ := c.Do(req)

	body, err := ioutil.ReadAll(rep.Body)

	log.Fatalln(string(body), err)
}
