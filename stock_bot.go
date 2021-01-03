package stock_telegram_bot

import (
	TgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var TgStockBot *TgBotApi.BotAPI

func InitTgStockBot() {
	bot, err := TgBotApi.NewBotAPI(GetConfig().Token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	TgStockBot = bot
}

func GetTgStockBot() *TgBotApi.BotAPI {
	if TgStockBot == nil {
		InitTgStockBot()
	}
	return TgStockBot
}
