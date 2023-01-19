package telegrambot

import (
	"fmt"
	"log"
	"medium_scraper/app"
	articleCommands "medium_scraper/app/article/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot      *tgbotapi.BotAPI
	Services app.Services
}

func NewBot(token string, services app.Services) Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	return Bot{
		Bot:      bot,
		Services: services,
	}
}

func (t Bot) ServeUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == `` || update.Message == nil {
			t.SendMessage(update.Message.Chat.ID, "invalid")
			continue
		}
		if !update.Message.IsCommand() {
			t.MessageHandler(*update.Message)
			continue
		}
		t.CommandHandler(*update.Message)

	}

}

func (t Bot) SendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)

	if _, err := t.Bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}

func (t Bot) MessageHandler(message tgbotapi.Message) {
	t.SendMessage(message.Chat.ID, "Im searching for that please wait ...")
	result, err := t.Services.ArticleService.Commands.SearchArticlesCommandHandler.Handle(
		articleCommands.SearchArticlesCommand{
			SearchText: message.Text,
		},
	)
	if err != nil {
		t.SendMessage(message.Chat.ID, "Something went wrong")
		return
	}
	for _, article := range result.Articles {
		articleMessage := (article.Title + "\n" + "\n" + article.Intro)
		t.SendMessage(message.Chat.ID, articleMessage)
	}
}

func (t Bot) CommandHandler(message tgbotapi.Message) {
	switch message.Command() {
	case "start":
		t.SendMessage(message.Chat.ID, welcomeMessage)
	default:
		t.SendMessage(message.Chat.ID, "Invalid command")
	}
}

var welcomeMessage string = ("Hi there \n" +
	"Welcome to Medium scraper telegram bot \n" +
	"This bot is still under development \n" +
	"But now, you can search medium for some articles \n" +
	"Type anything you want and ill search medium for it")
