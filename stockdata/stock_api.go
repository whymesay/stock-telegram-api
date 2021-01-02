package stockdata

type StockData struct {
	//名称
	Name string
	//代码
	Code string
	//涨幅
	Increase float32
	//当前价格
	CurrentPrice float32
	//昨天收盘价格
	YesterdayClosePrice float32
	//今天开盘价
	TodayOpenPrice float32
}

type StockApi interface {
	Query(code string) *StockData
	BatchQuery(codes []string) []*StockData
}
