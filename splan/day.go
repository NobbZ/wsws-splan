package splan

import (
	"strconv"

	"github.com/NobbZ/wsws-splan/settings"
)

type Days struct {
	Days []Day `json:"days"`
}

type Day struct {
	URI       string
	LongName  string
	ShortName string
}

func GetAllDays() Days {
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
