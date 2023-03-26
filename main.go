package main

import (
	"log"
	"os"

	"github.com/Ashkan0026/telegram-bot1/handlers"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	bot, err := tgbot.NewBotAPI(os.Getenv("token"))
	if err != nil {
		log.Println("Error while authenticating to the bot")
	}
	bot.Debug = false

	logger := log.Logger{}
	logger.SetOutput(os.Stdout)

	u := tgbot.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		logger.Printf("FirstName : %s, UserName : %s\n", update.Message.From.FirstName, update.Message.From.UserName)
		if update.Message.SenderChat != nil {
			logger.Printf("Tittle : %s, ID : %d", update.Message.SenderChat.Title, update.Message.SenderChat.ID)
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg := handlers.HandleHelp(update)
				_, err = bot.Send(msg)
				if err != nil {
					log.Printf("Error %v happened while sending the message\n", err)
				} else {
					log.Println("Message sent to chat ", update.Message.Chat.UserName)
				}

			case "start":
				msg := handlers.HandleStart(update)
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("%v happened while sending the message\n", err)
				}
			}
		} else {
			handlers.HandleMsg(update, bot)
		}
	}
}
