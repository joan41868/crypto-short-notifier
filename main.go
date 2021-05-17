package main

import (
	"crypto-short-notifier/service"
)

func main() {
	mp := make(map[string][]string)
	mp["https://api.binance.com/api/v3/ticker/24hr?symbol="] = []string{
		"ADAUSDT",
		"AAVEUSDT",
		"ETHUSDT",
		"BNBUSDT",
		"XRPUSDT",
		"XLMUSDT",
		"EOSUSDT",
		"TRXUSDT",
		"BCHUSDT",
		"DOTUSDT",
		"LTCUSDT",
		"FILUSDT",
		"UNIUSDT",
		"XTZUSDT",
		"YFIUSDT",
		"1INCHUSDT",
		"SUSHIUSDT",
	}
	collector := service.GetCollector(mp)
	collector.StartCollecting()
}
