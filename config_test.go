package stock_telegram_bot

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("./config.json")
	print(GetConfig().Token)
}
