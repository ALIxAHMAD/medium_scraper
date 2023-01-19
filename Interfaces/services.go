package interfaces

import (
	telegrambot "medium_scraper/Interfaces/telegram_bot"
	"medium_scraper/app"
)

type Services struct {
	Bot telegrambot.Bot
}

func NewServices(services app.Services, botToken string) Services {
	return Services{
		Bot: telegrambot.NewBot(botToken, services),
	}
}
