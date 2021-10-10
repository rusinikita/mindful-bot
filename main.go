package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rusinikita/discipline-bot/bot"
	"github.com/rusinikita/discipline-bot/db/airtable"
	"github.com/rusinikita/discipline-bot/task"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not loaded")
	}

	db, err := airtable.New()
	if err != nil {
		log.Fatal(err)

		return
	}

	b, err := telebot.NewBot(settings())
	if err != nil {
		log.Fatal(err)

		return
	}

	bot.RegisterHandlers(b, db, task.Handlers())

	b.Start()
}

func settings() telebot.Settings {
	return telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
}