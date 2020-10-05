package lazyfit

//TODO modularizzare
//TODO telegram
//TODO mini cache
//TODO db + login system

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"time"
)

func UnmarshalStatus(data []byte) (Status, error) {
	var r Status
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Status) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Status struct {
	Status  string  `json:"status"`
	Message *string `json:"message,omitempty"`
}

const (
	NONE Action = iota
	SCHEDULE
	BOOKING
	UNBOOK
	BOOKEDCOURSES
)

type Action uint8

type Config struct {
	Account  Account  `yaml:"account"`
	API      API      `yaml:"api"`
	Telegram Telegram `yaml:"telegram"`
}

var (
	start         = "2020-10-05"
	end           = "2020-10-12"
	currentDay    = time.Now() //.AddDate(0,0,-6) //TODO DEBUGONLY
	listaOrari    TimeTable
	prenotazione  Status
	listCourses   = make([]Course, 0, 35)
	listSchedules Schedules
	c             *colly.Collector
	action        Action
	Conf          *Config
	user          *Account
	api           *API
	teleBot       *Telegram
)

func Start() {
	user = NewAccount()
	api = NewApi()
	teleBot = NewTelegramBot()
	teleBot.Start()

	fmt.Println(currentDay.Format("Mon, 02/01/2006 15:04:05"))

	c = colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"),
	)

	c.OnResponse(func(r *colly.Response) {
		var err error

		switch action {
		case SCHEDULE:
			listaOrari, err = UnmarshalTimeTable(r.Body)
			if err != nil {
				log.Fatal(err)
			}
		case BOOKING:
			prenotazione, err = UnmarshalStatus(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(*prenotazione.Message)
		case UNBOOK:
			pre, err := UnmarshalStatus(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Esito cancellazione: ", pre.Status)
		case BOOKEDCOURSES:
			listSchedules, err = UnmarshalImpegni(r.Body)
			if err != nil {
				log.Fatal(err)
			}
		default:

		}
	})

	PreBookingInit()
	AutoBooking()
	teleBot.Stop()
}

func PreBookingInit() {
	user.Login()
	getCoursesRequest()
}

func AutoBooking() {
	giorno := currentDay.Weekday()
	start, end = getDailyFilterParam() //getWeeklyFilterParam()
	//giorno = time.Wednesday //TODO DEBUG ONLY
	switch giorno {
	case time.Monday:
		fallthrough
	case time.Wednesday:
		//Book("SALA PESI 17:00")//TODO DEBUG ONLY
		Book("CALISTHENICS")
		Book("SALA PESI 18:15")
		Book("SALA PESI 19:30")
	case time.Friday:
		Book("CALISTHENICS")
		Book("SALA PESI 17:00")
		Book("SALA PESI 19:30")
	}
}
