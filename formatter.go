package stock_telegram_bot

type Formatter interface {
	Format(s *StockData) string
}

type SimpleFormatter struct {
}

func (f *SimpleFormatter) Format(s *StockData) string {
	return s.Name + "|" + s.Increase + "%" + "\n"
}
