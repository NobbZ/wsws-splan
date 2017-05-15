package plansrc

import (
	"bytes"
	"encoding/xml"
	"strings"
	"testing"
)

func TestReadingMetaDataFromString(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected MetaData
	}{
		{"uhrzeit", "<erstellung><uhrzeit>10:55</uhrzeit></erstellung>", MetaData{Time: "10:55"}},
		{"datum", "<erstellung><datum>11.04.2014</datum></erstellung>", MetaData{Date: "11.04.2014"}},
		{"version", "<erstellung><version>2.2</version></erstellung>", MetaData{Version: "2.2"}},
		{"proginfo", "<erstellung><proginfo>2.2</proginfo></erstellung>", MetaData{ProgInfo: "2.2"}},
		{"hochschulsport", "<erstellung><hochschulsport>2.2</hochschulsport></erstellung>", MetaData{Sport: "2.2"}},
		{"Autor", "<erstellung><autor>Birger Wolter</autor></erstellung>", MetaData{Author: "Birger Wolter"}},
		{"Semestertyp (Sommer)", "<erstellung><semestertyp>SS</semestertyp></erstellung>", MetaData{SemType: Summer}},
		{"Semestertyp (Winter)", "<erstellung><semestertyp>WS</semestertyp></erstellung>", MetaData{SemType: Winter}},
	}

	read := func(in string) MetaData {
		meta := MetaData{}

		r := strings.NewReader(in)

		xml.NewDecoder(r).Decode(&meta)

		return meta
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := read(test.input)
			if actual != test.expected {
				t.Errorf("Expected: %#v;\nGot: %#v", test.expected, actual)
			}
		})
	}
}

func TestReadingDayFromString(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected Day
	}{}

	read := func(in string) Day {
		day := Day{}

		r := strings.NewReader(in)

		xml.NewDecoder(r).Decode(&day)

		return day
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := read(test.input)
			if actual != test.expected {
				t.Errorf("Expected: %#v;\nGot: %#v", test.expected, actual)
			}
		})
	}
}

func TestMarshalSemType(t *testing.T) {
	buf := new(bytes.Buffer)

	xml.NewEncoder(buf).Encode(MetaData{SemType: Summer})

	if buf.String() != "<erstellung><uhrzeit></uhrzeit><datum></datum><version></version><proginfo></proginfo><hochschulsport></hochschulsport><autor></autor><string>SS</string></erstellung>" {
		t.Error(buf.String())
	}

	buf = new(bytes.Buffer)
	xml.NewEncoder(buf).Encode(MetaData{SemType: Winter})
	if buf.String() != "<erstellung><uhrzeit></uhrzeit><datum></datum><version></version><proginfo></proginfo><hochschulsport></hochschulsport><autor></autor><string>WS</string></erstellung>" {
		t.Error(buf.String())
	}
}

func TestUnmarshalSemType(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected SemType
	}{
		{"Sommer", "<semestertyp>SS</semestertyp>", Summer},
		{"Winter", "<semestertyp>WS</semestertyp>", Winter},
	}

	for _, test := range testData {
		t.Run(test.name, func(t *testing.T) {
			var actual SemType
			r := strings.NewReader(test.input)
			xml.NewDecoder(r).Decode(&actual)

			if actual != test.expected {
				t.Error("Expected %v, but got %v", test.expected, actual)
			}
		})
	}
}
