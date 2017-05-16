package splan

type Days struct {
	Days []Day `json:"days"`
}

type Day struct {
	ID        int
	LongName  string
	ShortName string
}

func GetAllDays() Days {
	ds := Days{
		Days: make([]Day, len(splanXML.Chitage.Chitag)),
	}
	for i, d := range splanXML.Chitage.Chitag {
		day := Day{
			ID:        int(d.Chiid[0].Text),
			ShortName: d.Chibezeichnung.Chikurz.Text,
			LongName:  d.Chibezeichnung.Chilang.Text,
		}
		ds.Days[i] = day
	}

	return ds
}
