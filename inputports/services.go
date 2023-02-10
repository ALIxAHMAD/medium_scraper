package interfaces

import (
	"medium_scraper/app"
	telegrambot "medium_scraper/inputports/telegram_bot/bot"
	botdatabase "medium_scraper/inputports/telegram_bot/bot_database"
)

type Services struct {
	Bot telegrambot.Bot
}

func NewServices(services app.Services, botToken string, db botdatabase.BotDataBase) Services {
	return Services{
		Bot: telegrambot.NewBot(botToken, services, db),
	}
}
