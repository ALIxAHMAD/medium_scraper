package telegrambot

import (
	"fmt"
	"log"
	botdatabase "medium_scraper/Interfaces/telegram_bot/bot_database"
	"medium_scraper/app"
	articleCommands "medium_scraper/app/article/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot      *tgbotapi.BotAPI
	Db       botdatabase.BotDataBase
	Services app.Services
}

func NewBot(token string, services app.Services, db botdatabase.BotDataBase) Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	return Bot{
		Bot:      bot,
		Db:       db,
		Services: services,
	}
}

func (t Bot) ServeUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message.Text == `` || update.Message == nil {
			t.sendMessage(update.Message.Chat.ID, "invalid")
			continue
		}
		if !update.Message.IsCommand() {
			t.messageHandler(*update.Message)
			continue
		}
		t.commandHandler(*update.Message)

	}

}

func (t Bot) messageHandler(message tgbotapi.Message) {
	chatId := message.Chat.ID
	isNew, err := t.newUserCheck(message)
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	if isNew {
		t.sendMessage(chatId, welcomeMessage)
	}
	user, err := t.Db.GetUser(chatId)
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	command := user.LastCommand
	switch command {
	case "search":
		t.searchMessageHandler(message)
	default:
		t.sendMessage(chatId, welcomeMessage)
	}
}

func (t Bot) sendMessage(chatId int64, message string) {
	msg := tgbotapi.NewMessage(chatId, message)

	if _, err := t.Bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}

func (t Bot) searchMessageHandler(message tgbotapi.Message) {
	chatId := message.Chat.ID
	command := ""
	err := t.Db.UpdateUser(chatId, command)
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	t.sendMessage(message.Chat.ID, "Im searching for that please wait ...")
	result, err := t.Services.ArticleService.Commands.SearchArticlesCommandHandler.Handle(
		articleCommands.SearchArticlesCommand{
			SearchText: message.Text,
		},
	)
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	for _, article := range result.Articles {
		articleMessage := (article.Title + "\n\n" + article.Intro)
		t.sendMessage(message.Chat.ID, articleMessage)
	}
}

func (t Bot) commandHandler(message tgbotapi.Message) {
	switch message.Command() {
	case "start":
		t.startCommandHandler(message)
	case "search":
		t.searchCommandHandler(message)
	default:
		t.sendMessage(message.Chat.ID, "Invalid command")
	}
}

func (t Bot) startCommandHandler(message tgbotapi.Message) {
	_, err := t.newUserCheck(message)
	if err != nil {
		t.sendErrorMessage(message.Chat.ID)
		return
	}
	t.sendMessage(message.Chat.ID, welcomeMessage)
}

func (t Bot) newUserCheck(message tgbotapi.Message) (bool, error) {
	chatId := message.Chat.ID
	isNew, err := t.Db.IsNewUser(chatId)
	if err != nil {
		t.sendErrorMessage(chatId)
		return isNew, err
	}

	if isNew {
		var user = botdatabase.User{
			FirstName:   message.Chat.FirstName,
			LastName:    message.Chat.LastName,
			LastCommand: "",
		}
		err = t.Db.AddUser(chatId, user)
		if err != nil {
			return isNew, err
		}
		return isNew, nil
	}
	return isNew, nil
}

func (t Bot) searchCommandHandler(message tgbotapi.Message) {
	chatId := message.Chat.ID
	_, err := t.newUserCheck(message)
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	err = t.Db.UpdateUser(chatId, "search")
	if err != nil {
		t.sendErrorMessage(chatId)
		return
	}
	t.sendMessage(chatId, "Enter search text:")
}

func (t Bot) sendErrorMessage(chatId int64) {
	t.sendMessage(chatId, "Something went wrong")
}

var welcomeMessage string = ("Hi there \n" +
	"Welcome to Medium scraper telegram bot \n" +
	"This bot is still under development \n\n" +
	"/search-search medium for some articles \n")
