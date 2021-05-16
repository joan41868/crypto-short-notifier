package main

import (
	"pw-be-dc/service"
)

/**
  Main idea - collect data from various sources.
  Should utilize Binance API

  Binance - periodically fetch shortable gainers:
  	1. Get coins which have leverages - use db, as more can come
  	2. Analyze percentage - if it's over 30% - push to fe
  	3. List of BLVT-able tokens:
  		3.1 ADA
  		3.2 AAVE
  		3.3 ETH
  		3.5 BNB
  		3.6 XRP
  		3.7 XLM
  		3.8 EOS
  		3.9 TRX
  		3.10

*/
func main() {

	var mp map[string][]string
	mp = make(map[string][]string)
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