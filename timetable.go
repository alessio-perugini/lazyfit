package lazyfit

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

const (
	LAYOUTZONE = "2006-01-02T15:04:05"
)

type TimeTable []TimeElement

func UnmarshalTimeTable(data []byte) (TimeTable, error) {
	var r TimeTable
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TimeTable) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TimeElement struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	AllDay        bool        `json:"allDay"`
	StartDateTime string      `json:"startDateTime"`
	EndDateTime   string      `json:"endDateTime"`
	URL           interface{} `json:"url"`
	ClassName     string      `json:"className"`
	Editable      bool        `json:"editable"`
	Start         string      `json:"start"`
	End           string      `json:"end"`
	ExtraData     ExtraData   `json:"extraData"`
}

type ExtraData struct {
}

func getTimeTableRequest(corso, startTime, endTime string) {
	action = SCHEDULE
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", api.NewBook)
	})

	err := c.Post(api.TimeTable, map[string]string{"idattivita": corso, "idlivello": "0", "start": startTime, "end": endTime})
	if err != nil {
		log.Fatal(err)
	}
}

func getTodayLastAvailableCourse() string {
	if nFasce := len(listaOrari); nFasce > 0 { /*
			t, err := time.Parse(LAYOUTZONE, listaOrari[nFasce-1].StartDateTime)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(t)*/
		return listaOrari[nFasce-1].ID
	} else {
		return ""
	}
}

func PrintInfoOrari() {
	for _, v := range listaOrari {
		fmt.Printf("%s %s %s\n", strings.ReplaceAll(v.Title, "\n", ""), v.StartDateTime, v.EndDateTime)
	}
}
