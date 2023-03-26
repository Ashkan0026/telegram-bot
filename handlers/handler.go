package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Ashkan0026/telegram-bot1/api"
	"github.com/Ashkan0026/telegram-bot1/models"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleHelp(update tgbot.Update) tgbot.MessageConfig {
	msg := tgbot.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "This bot helps you to get information abount cryptocurrencies\n" +
		"type coin + coinName to get information about specified coin\n" +
		"type coin + allCoins to get info about top 20 coins"
	return msg
}

func HandleStart(update tgbot.Update) tgbot.MessageConfig {
	msg := tgbot.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hello From David the Holy\nType /help to get help"
	return msg
}

func HandleMsg(update tgbot.Update, bot *tgbot.BotAPI) {
	msg := update.Message.Text
	msg = strings.ToLower(msg)
	if strings.Contains(msg, "coin") {
		msg = strings.TrimSpace(msg)
		data := string(msg[4:])
		data = strings.TrimSpace(data)
		HandleCoin(update, bot, data)
	} else if msg == "today" {
		date := time.Now()
		monthStr := ""
		if int(date.Month()) < 10 {
			monthStr = "0" + strconv.Itoa(int(date.Month()))
		} else {
			monthStr = strconv.Itoa(int(date.Month()))
		}
		dayStr := ""
		if date.Day() < 10 {
			dayStr = "0" + strconv.Itoa(date.Day())
		} else {
			dayStr = strconv.Itoa(date.Day())
		}
		formattedDate := strconv.Itoa(date.Year()) + "-" + monthStr + "-" + dayStr
		msgTel := tgbot.NewMessage(update.Message.Chat.ID, formattedDate)
		SendMessage(bot, msgTel)
	}
}

func HandleCoin(update tgbot.Update, bot *tgbot.BotAPI, txt string) {
	msg := tgbot.NewMessage(update.Message.Chat.ID, "")
	fmt.Println(txt)
	if txt != "allcoins" {
		coinId := strings.ToLower(txt)
		coin := api.GetCurrency(coinId)
		if coin == nil {
			msg.Text = "Sorry try Later"
		} else if coin.Coin == nil {
			msg.Text = "There isn't such coin"
		} else {
			msg.Text = coin.Coin.String()
		}
		SendMessage(bot, msg)
	} else {
		coins := api.GetCurrencies()
		if coins == nil {
			msg.Text = "Failed to get data"
			SendMessage(bot, msg)
		} else {
			SendMultipleMessages(bot, update, coins)
		}
	}
}

func SendMessage(bot *tgbot.BotAPI, msg tgbot.MessageConfig) {
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Msg wasn't send\n")
	}
}

func SendMultipleMessages(bot *tgbot.BotAPI, update tgbot.Update, coins []*models.Coin) {
	msg := tgbot.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Wolf of wallstreet Or Abu sofian ?"
	SendMessage(bot, msg)
	for _, coin := range coins {
		coinMsg := tgbot.NewMessage(update.Message.Chat.ID, "")
		coinMsg.Text = coin.String()
		SendMessage(bot, coinMsg)
	}
}
