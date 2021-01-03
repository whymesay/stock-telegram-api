package stock_telegram_bot

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("./config.json")
	fmt.Println(Token)
}
