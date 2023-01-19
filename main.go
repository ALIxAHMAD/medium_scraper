package main

import (
	"log"
	"medium_scraper/Infrastructure/article/scraper"
	interfaces "medium_scraper/Interfaces"
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

	appServices := app.NewServices(&repo)

	interfacesServices := interfaces.NewServices(appServices, config.BotToken)

	interfacesServices.Bot.ServeUpdates()
}
