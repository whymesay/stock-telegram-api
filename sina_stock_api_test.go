package stock_telegram_bot

import (
	"fmt"
	"testing"
)

func TestBatchQuery(t *testing.T) {

	sinaStockApi := &SinaStockApi{}
	var param = []string{"sh000001", "sh000002"}
	result := sinaStockApi.BatchQuery(param)
	for _, data := range result {
		fmt.Println(data)
	}
}
