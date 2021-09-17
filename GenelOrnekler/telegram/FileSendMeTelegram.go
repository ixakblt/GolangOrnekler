package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, _ := tgbotapi.NewBotAPI("149825515asdwKk_kh5iqyiphQ_cjiCVTz0n9c9HZ24")
	bot.Debug = true
	var ChatID int64
	ChatID = 1316224975 //Chatid

	for i := 1; i < len(os.Args); i++ {
		msg := tgbotapi.NewDocumentUpload(ChatID, os.Args[i])
		bot.Send(msg)
	}

}
