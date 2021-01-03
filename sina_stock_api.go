package stock_telegram_bot

import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const SinaStockUrl = "http://hq.sinajs.cn/list="

type SinaStockApi struct {
}

func (sinaStockApi *SinaStockApi) Query(code string) *StockData {
	return sinaStockApi.BatchQuery([]string{code})[0]
}

func (sinaStockApi *SinaStockApi) BatchQuery(codes []string) []*StockData {
	var param = strings.Join(codes, ",")
	resp, err := http.Get(SinaStockUrl + param)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	deCodeBody := strings.Trim(mahonia.NewDecoder("gbk").ConvertString(string(body)), " ")

	lineList := strings.Split(deCodeBody, ";\n")
	var stockDataList []*StockData
	for _, line := range lineList {
		if line != "" {
			stockData := convertStockData(line)
			stockDataList = append(stockDataList, stockData)
		}
	}
	return stockDataList
}

func convertStockData(line string) *StockData {
	content := strings.Split(line[strings.Index(line, "\""):strings.LastIndex(line, ",\"")], ",")

	todayOpenPrice, _ := strconv.ParseFloat(content[1], 64)
	yesterdayClosePrice, _ := strconv.ParseFloat(content[2], 64)
	currentPrice, _ := strconv.ParseFloat(content[3], 64)

	code := line[11:19]
	stockData := &StockData{
		Name:                content[0][1:],
		Code:                code,
		Increase:            strconv.FormatFloat(float64((float32(currentPrice)-float32(yesterdayClosePrice))/float32(yesterdayClosePrice)*100), 'f', 2, 64),
		CurrentPrice:        strconv.FormatFloat(currentPrice, 'f', 2, 64),
		YesterdayClosePrice: strconv.FormatFloat(yesterdayClosePrice, 'f', 2, 64),
		TodayOpenPrice:      strconv.FormatFloat(todayOpenPrice, 'f', 2, 64),
	}
	return stockData
}
