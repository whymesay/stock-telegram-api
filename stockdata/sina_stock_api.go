package stockdata

import (
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
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

func (s StockData) String() string {
	return s.Name + s.Code
}

func convertStockData(line string) *StockData {
	content := strings.Split(line[strings.Index(line, "\""):strings.LastIndex(line, ",\"")], ",")
	code := line[11:19]
	stockData := &StockData{
		Name:                content[0][1:],
		Code:                code,
		Increase:            0,
		CurrentPrice:        0,
		YesterdayClosePrice: 0,
		TodayOpenPrice:      0,
	}
	return stockData
}
