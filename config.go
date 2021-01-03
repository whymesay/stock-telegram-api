package stock_telegram_bot

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type BotConfig struct {
	Token        string
	StockDataApi StockApi
}

var Config *BotConfig

func LoadConfig(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json") //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	Config = &BotConfig{
		Token:        viper.GetString("token"),
		StockDataApi: ApiFactory(viper.GetString("apiSource")),
	}
}

func GetConfig() *BotConfig {
	return Config
}
