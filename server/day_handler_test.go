package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
)

func TestDayOptions(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go Start("localhost:9191", wg)
	wg.Wait()

	req, _ := http.NewRequest("GET", "http://localhost:9191/day", nil)
	c := new(http.Client)
	req.Header["Accept"] = []string{"application/xml"}
	rep, _ := c.Do(req)

	body, err := ioutil.ReadAll(rep.Body)

	log.Fatalln(string(body), err)
}
