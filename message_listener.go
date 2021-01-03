package stock_telegram_bot

import (
	TgBotApi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func ListenMessage() {
	bot := GetTgStockBot()
	u := TgBotApi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		resolveMessage(update, bot)
	}

}

func resolveMessage(update TgBotApi.Update, bot *TgBotApi.BotAPI) {
	msg := TgBotApi.NewMessage(update.Message.Chat.ID, ExecCommand(update.Message.Text))
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Reply Message Fail", err)
	}
}
