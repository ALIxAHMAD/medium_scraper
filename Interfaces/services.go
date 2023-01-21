package interfaces

import (
	telegrambot "medium_scraper/Interfaces/telegram_bot/bot"
	botdatabase "medium_scraper/Interfaces/telegram_bot/bot_database"
	"medium_scraper/app"
)

type Services struct {
	Bot telegrambot.Bot
}

func NewServices(services app.Services, botToken string, db botdatabase.BotDataBase) Services {
	return Services{
		Bot: telegrambot.NewBot(botToken, services, db),
	}
}
