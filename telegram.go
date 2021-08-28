package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Telegram() {
	b, err := tb.NewBot(tb.Settings{
		Token:  GetConfig().TelegramBotAPIKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
}
