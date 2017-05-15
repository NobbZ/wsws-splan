package server

import (
	"log"
	"net/http"
	"testing"
)

func TestDayOptions(t *testing.T) {
	go Start("localhost:9191")

	req, _ := http.NewRequest("OPTIONS", "http://localhost:9191/day/2", nil)
	c := new(http.Client)
	rep, err := c.Do(req)

	log.Fatalln(rep, err)
}
