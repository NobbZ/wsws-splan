package splan

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

var splanXML = new(Chisplan)

var PathToXML = "/home/nmelzer/gopath/src/github.com/NobbZ/wsws-splan"

func init() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("An error occured while opening %v:\n%v", filename, err)
	}

	xmlData, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("An error occured while reading %v:\n%v", filename, err)
	}

	err = xml.Unmarshal(xmlData, splanXML)
	if err != nil {
		log.Fatalf("An error occured while processing %v:\n%v", filename, err)
	}
}

type SPlan struct{}
