package main

import (
	"log"
	"medium_scraper/Infrastructure/articles/scraper"
	telegrambot "medium_scraper/Interfaces/telegram_bot"
	"medium_scraper/util/env"
)

func main() {
	env.Load(".env")

	config, err := env.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	repo := scraper.NewRepo()

	bot := telegrambot.NewTelegramBot(config.BotToken, &repo)

	bot.GetUpdates()
}
