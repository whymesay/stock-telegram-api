package stock_telegram_bot

func init() {
	LoadConfig("./myconfig.json")
	InitTgStockBot()
}
