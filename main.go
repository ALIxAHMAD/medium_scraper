package main

import (
	"fmt"
	"log"
	"medium_scraper/app"
	interfaces "medium_scraper/inputports"
	redisdb "medium_scraper/inputports/telegram_bot/bot_database/redis"
	"medium_scraper/interfaceadapters/article/scraper"
	"medium_scraper/util/env"
)

func main() {
	env.Load(".env")

	config, err := env.ParseConfig()

	if err != nil {
		log.Fatal(err)
	}

	repo := scraper.NewRepo()

	db := redisdb.NewRepo(config.RedisUrl, config.RedisPassword)
	fmt.Println("connected to database")

	appServices := app.NewServices(&repo)

	interfacesServices := interfaces.NewServices(appServices, config.BotToken, db)

	interfacesServices.Bot.ServeUpdates()
}
