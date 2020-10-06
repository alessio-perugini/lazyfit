package lazyfit

import (
	"github.com/gocolly/colly"
	"log"
)

func bookRequest(idCorso, idImpegno string) {
	action = BOOKING
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", api.NewBook)
	})

	err := c.Post(api.Book, map[string]string{"idattivita": idCorso, "idimpegno": idImpegno, "libera": "false"})
	if err != nil {
		log.Fatal(err)
	}
}

func unBookRequest(req map[string]string) {
	action = UNBOOK
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", api.BaseBooking)
	})

	err := c.Post(api.Delete, req)
	if err != nil {
		log.Fatal(err)
	}
}

func Book(nomeCorso string) {
	idCourse := findCourseIdFromName(nomeCorso)
	getTimeTableRequest(idCourse, start, end)
	idBooking := getTodayLastAvailableCourse()
	if idBooking == "" {
		teleBot.SendTelegramMessage(nomeCorso + ": corso pieno")
		return
	}

	bookRequest(idCourse, idBooking)
	teleBot.SendTelegramMessage(nomeCorso + ": " + *prenotazione.Message)
}

//TODO migliorare questo xke dovrei fare un fetch dei prenotati e poi da li scegliere quale
func UnBook(idBooked string) {
	getBookedCoursesRequest()
	info := getBookedCourseInfo(idBooked)
	if info.Idimpegno == "" {
		log.Println("ERRORE non è stato possibile togliere la prenotazione")
		return
	}
	queryParam := ConvertStructToMapOfStrings(info)
	unBookRequest(queryParam)
}
