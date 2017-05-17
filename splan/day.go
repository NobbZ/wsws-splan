package splan

import (
	"strconv"

	"github.com/NobbZ/wsws-splan/settings"
)

type Days struct {
	XMLName interface{} `json:",omitempty" xml:"days"`
	Days    []Day       `json:"days"       xml:"day"`
}

type Day struct {
	URI       string `json:"uri"        xml:"uri"`
	LongName  string `json:"long_name"  xml:"long_name"`
	ShortName string `json:"short_name" xml:"short_name"`
}

func GetAllDays(depth int) Days {
	ds := Days{
		Days: make([]Day, len(splanXML.Chitage.Chitag)),
	}
	for i, d := range splanXML.Chitage.Chitag {
		day := Day{
			URI:       settings.BaseURI() + "day/" + strconv.Itoa(int(d.Chiid[0].Text)),
			ShortName: d.Chibezeichnung.Chikurz.Text,
			LongName:  d.Chibezeichnung.Chilang.Text,
		}
		ds.Days[i] = day
	}

	return ds
}
