package splan

import "strconv"

type Days struct {
	Days []Day `json:"days"`
}

type Day struct {
	URI       string
	LongName  string
	ShortName string
}

func GetAllDays(path string) Days {
	ds := Days{
		Days: make([]Day, len(splanXML.Chitage.Chitag)),
	}
	for i, d := range splanXML.Chitage.Chitag {
		day := Day{
			URI:       path + strconv.Itoa(int(d.Chiid[0].Text)),
			ShortName: d.Chibezeichnung.Chikurz.Text,
			LongName:  d.Chibezeichnung.Chilang.Text,
		}
		ds.Days[i] = day
	}

	return ds
}
