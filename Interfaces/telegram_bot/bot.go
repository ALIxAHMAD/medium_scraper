package telegrambot

import (
	"medium_scraper/domain/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	Bot  *tgbotapi.BotAPI
	Repo repository.ArticleRepository
}

func NewTelegramBot(token string, repo repository.ArticleRepository) TelegramBot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	return TelegramBot{
		Bot:  bot,
		Repo: repo,
	}
}

func (t TelegramBot) GetUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		t.SendMessage(update.Message.Chat.ID, "Im searching for that please wait ...")
		articles, err := t.Repo.SearchArticles(update.Message.Text)
		if err != nil {
			t.SendMessage(update.Message.Chat.ID, "Something went wrong")
		}
		for _, article := range articles {
			articleMessage := (article.Title + "\n" + "\n" + article.Intro)
			t.SendMessage(update.Message.Chat.ID, articleMessage)
		}
	}

}

func (t TelegramBot) SendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)

	if _, err := t.Bot.Send(msg); err != nil {
		panic(err)
	}
}
