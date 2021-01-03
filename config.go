package stock_telegram_bot

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Token string

func LoadConfig(configFile string) {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json") //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	Token = viper.GetString("token")
}
