package plansrc

import "os"
import "encoding/xml"
import "strings"
import "errors"

type SemType byte

const (
	Summer SemType = iota
	Winter
)

func (st SemType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	switch st {
	case Summer:
		return e.Encode("SS")
	case Winter:
		return e.Encode("WS")
	default:
		return errors.New("Invalid SemType")
	}
}

func (st *SemType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	token, err := d.Token()
	if err != nil {
		return err
	}

	switch t := token.(type) {
	case xml.CharData:
		switch string(t) {
		case "SS":
			*st = Summer
		case "WS":
			*st = Winter
		}
	}

	return nil
}

type Splan struct {
	XMLName struct{} `xml:"splan"`
	Meta    MetaData `xml:"erstellung"`
	Days    Days     `xml:"tage"`
}

type MetaData struct {
	XMLName  struct{} `xml:"erstellung"`
	Time     string   `xml:"uhrzeit"`
	Date     string   `xml:"datum"`
	Version  string   `xml:"version"`
	ProgInfo string   `xml:"proginfo"`
	Sport    string   `xml:"hochschulsport"`
	Author   string   `xml:"autor"`
	SemType  SemType  `xml:"semestertyp"`
}

type Days struct {
	XMLName struct{} `xml:"tage"`
	ID2Name map[byte]Day
}

func (ds *Days) UnmarshalXML(UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

}

type Day struct {
	XMLName struct{} `xml:"tag"`
	Short   string
	Long    string
}

func FromFile(path string) (Splan, error) {
	splan := Splan{}

	r, err := os.Open(path)
	if err != nil {
		return Splan{}, err
	}

	xml.NewDecoder(r).Decode(&splan)

	return splan, nil
}

func FromString(input string) (Splan, error) {
	splan := Splan{}

	r := strings.NewReader(input)

	xml.NewDecoder(r).Decode(&splan)

	return splan, nil
}
