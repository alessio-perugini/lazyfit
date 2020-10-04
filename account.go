package lazyfit

import (
	"github.com/gocolly/colly"
	"log"
)

type Account struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewAccount() *Account{
	return &Conf.Account
}

func (a *Account) Login(){
	loginRequest(a.Username, a.Password)
}

func loginRequest(username, password string) {
	action = NONE

	c.OnHTML("form.login-form", func(e *colly.HTMLElement) {
		token := e.ChildAttr("input[name='__RequestVerificationToken']", "value")
		action = NONE

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Referer", api.Login)
		})

		err := c.Post(api.Login, map[string]string{"__RequestVerificationToken": token, "UserName": username, "Password": password, "RememberMe": "false"})
		if err != nil {
			log.Fatal(err)
		}
	})

	err := c.Visit(api.Login)
	if err != nil {
		log.Fatal(err)
	}
}