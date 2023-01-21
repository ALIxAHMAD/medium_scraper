package main

import (
	"fmt"
	"log"
	"medium_scraper/Infrastructure/article/scraper"
	interfaces "medium_scraper/Interfaces"
	redisdb "medium_scraper/Interfaces/telegram_bot/bot_database/redis"
	"medium_scraper/app"
	"medium_scraper/util/env"
)

func main() {
	env.Load(".env")

	config, err := env.ParseConfig()

	if err != nil {
		log.Fatal(err)
	}

	repo := scraper.NewRepo()

	db := redisdb.NewRepo(config.RedisUrl, "")
	fmt.Println("connected to database")

	appServices := app.NewServices(&repo)

	interfacesServices := interfaces.NewServices(appServices, config.BotToken, db)

	interfacesServices.Bot.ServeUpdates()
}
