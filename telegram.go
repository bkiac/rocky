package main

import (
	"fmt"
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Telegram() {
	b, err := tb.NewBot(tb.Settings{
		Token:  TelegramBotAPIKey,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		err := b.Notify(m.Sender, "typing")
		if err != nil {
			log.Default().Println(err)
			return
		}

		if IMDbURLRegexp.MatchString(m.Text) {
			mos, err := GetMovieOrSeries(m.Text)
			if err != nil {
				_, err := b.Send(m.Sender, err)
				if err != nil {
					log.Default().Println(err)
				}
				return
			}

			roam, err := MovieOrSeriesToRoamPage(mos)
			if err != nil {
				_, err := b.Send(m.Sender, err)
				if err != nil {
					log.Default().Println(err)
				}
				return
			}

			_, err = b.Send(m.Sender, roam)
			if err != nil {
				log.Default().Println(err)
			}
			return
		}

		if GoodreadsURLRegexp.MatchString(m.Text) {
			book, err := GetBook(m.Text)
			if err != nil {
				_, err := b.Send(m.Sender, err)
				if err != nil {
					log.Default().Println(err)
				}
				return
			}

			roam := BookToRoamPage(book, false)
			_, err = b.Send(m.Sender, roam)
			if err != nil {
				log.Default().Println(err)
			}
			return
		}

		_, err = b.Send(m.Sender, "Invalid Message")
		if err != nil {
			log.Default().Println(err)
		}
	})

	fmt.Println("Starting Telegram Bot...")
	b.Start()
}
