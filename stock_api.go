package stock_telegram_bot

import "encoding/json"

type StockData struct {
	//名称
	Name string
	//代码
	Code string
	//涨幅
	Increase string
	//当前价格
	CurrentPrice string
	//昨天收盘价格
	YesterdayClosePrice string
	//今天开盘价
	TodayOpenPrice string
}

func (s StockData) String() string {
	marshal, _ := json.Marshal(s)
	return string(marshal)
}

type StockApi interface {
	Query(code string) *StockData
	BatchQuery(codes []string) []*StockData
}

func ApiFactory(apiSource string) StockApi {
	if apiSource == "sina" {
		return new(SinaStockApi)
	}
	panic("must have api source ")
	return nil
}
