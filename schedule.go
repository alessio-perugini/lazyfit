package lazyfit

import (
	"encoding/json"
	"fmt"
)

func UnmarshalImpegni(data []byte) (Schedules, error) {
	var r Schedules
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Schedules) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Schedules struct {
	Data []Schedule `json:"data"`
}

type Schedule struct {
	Idimpegno       string      `json:"idimpegno"`
	Data            string      `json:"data"`
	DataoraSort     string      `json:"dataoraSort"`
	Orai            string      `json:"orai"`
	Oraf            string      `json:"oraf"`
	Note            string      `json:"note"`
	Attivita        string      `json:"attivita"`
	Livello         interface{} `json:"livello"`
	Idoggetto       string      `json:"idoggetto"`
	Idnoleggio      string      `json:"idnoleggio"`
	Idattivita      string      `json:"idattivita"`
	Idattesa        string      `json:"idattesa"`
	IsPrenotazione  bool        `json:"isPrenotazione"`
	CanDelete       bool        `json:"canDelete"`
	CanDeleteAttesa bool        `json:"canDeleteAttesa"`
	CanAssenza      bool        `json:"canAssenza"`
	IsAssente       bool        `json:"isAssente"`
	Tipooggetto     string      `json:"tipooggetto"`
	Tipo            string      `json:"tipo"`
	Descrizione     string      `json:"descrizione"`
	Sposta          bool        `json:"sposta"`
	Giorno          string      `json:"giorno"`
	DTRowClass      string      `json:"DT_RowClass"`
	ActionButtons   string      `json:"ActionButtons"`
}

func getBookedCourseInfo(idCourse string) Schedule {
	bookedList := listSchedules.Data
	for i := range bookedList {
		if bookedList[i].Idimpegno == idCourse {
			return bookedList[i]
		}
	}

	return Schedule{}
}

func getBookedCoursesRequest() {
	action = BOOKEDCOURSES

	err := c.Visit(api.Schedules)
	if err != nil {
		fmt.Println(err)
	}
}