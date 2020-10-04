package lazyfit

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Course struct {
	nome string
	id   string
}

func getCoursesRequest() {
	//Get lezioni
	action = NONE
	c.OnHTML("div.list-group", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, element *colly.HTMLElement) {
			id := element.Attr("data-idattivita")
			nome := element.ChildText("div h3")
			corso := Course{nome: nome, id: id}
			listCourses = append(listCourses, corso)
		})
	})

	err := c.Visit(api.Courses)
	if err != nil {
		fmt.Println(err)
	}
}

func PrintCourses(){
	for _, v := range listCourses {
		fmt.Printf("nome: %s \t id: %s\n", v.nome, v.id)
	}
}