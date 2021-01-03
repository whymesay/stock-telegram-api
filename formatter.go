package stock_telegram_bot

import (
	"encoding/json"
	"fmt"
	"github.com/whymesay/stock-telegram-bot/stockdata"
)

type FormatData stockdata.StockData

type Formatter interface {
	Decorator(s *FormatData) *FormatData
	Format(s FormatData) string
}

type AbstractFormatter struct {
}

func (f AbstractFormatter) Format(s FormatData) string {
	marshal, _ := json.Marshal(s)
	return string(marshal)
}
func (f AbstractFormatter) Decorator(s *FormatData) *FormatData {
	sprintf := fmt.Sprintf("%v", s)
	panic("not support decorator" + sprintf)
}

type SimpleFormatter struct {
	AbstractFormatter
}

func (f SimpleFormatter) Decorator(s *FormatData) *FormatData {
	return &FormatData{
		Name: resolveValue(s.Name, func() string {
			return s.Name
		}),
		Code: resolveValue(s.Code, func() string {
			return ""
		}),
		Increase:            "",
		CurrentPrice:        "",
		YesterdayClosePrice: "",
		TodayOpenPrice:      "",
	}
}

type ColorFormatter struct {
}

func (f ColorFormatter) Decorator(s *FormatData) *FormatData {
	return &FormatData{
		Name: resolveValue(s.Name, func() string {
			return "<span style='color:red'>" + s.Name + "<span>"
		}),
		Code:                "",
		Increase:            "",
		CurrentPrice:        "",
		YesterdayClosePrice: "",
		TodayOpenPrice:      "",
	}
}

func resolveValue(value string, fn func() string) string {
	if value != "" {
		return fn()
	}
	return ""
}
