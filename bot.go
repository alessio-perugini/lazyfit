package lazyfit

import (
	tb "github.com/tucnak/telebot"
	"log"
	"time"
)

type Telegram struct {
	Token  string `yaml:"token"`
	Chatid string `yaml:"chatid"`
	tBot   *tb.Bot
}

func NewTelegramBot() *Telegram {
	return &Conf.Telegram
}

func (t *Telegram) Start() {
	var err error
	t.tBot, err = tb.NewBot(tb.Settings{
		Token:  t.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	t.tBot.Handle("/hello", func(m *tb.Message) {
		t.tBot.Send(m.Sender, "Hello World!")
	})

	go t.tBot.Start()
}

func (t *Telegram) SendTelegramMessage(text string) {
	chat, _ := t.tBot.ChatByID(t.Chatid)
	t.tBot.Send(chat, text)
}

func (t *Telegram) Stop() {
	t.tBot.Stop()
}
