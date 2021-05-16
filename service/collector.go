package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type CoinData struct {
	Symbol string `json:"symbol"`
	PriceChangePercent string `json:"priceChangePercent"`
	Volume string `json:"volume"`
}

type Collector struct {
	// Url -> [] endpoints
	UrlToEndpointsMap map[string] []string
	urls []string
	Client *http.Client
}

func (c *Collector) StartCollecting(){
	for {
		var data []*CoinData
		for _, url := range c.urls{
			resp, err := c.Client.Get(url)

			if err != nil {
				log.Println(err)
			}

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Println(err)
			}

			var coinData *CoinData
			err = json.Unmarshal(body, &coinData)

			if err != nil{
				log.Println(err)
			}
			//log.Println(coinData)
			data = append(data, coinData)
		}

		for _, c := range data{
			percentage, err  := strconv.ParseFloat(c.PriceChangePercent, 2)
			if err != nil{
				log.Println(err)
			}
			if percentage > 15 {
				// send notifications
				msg := fmt.Sprintf("%s is %f UP. Consider buying %s-DOWN token.")
				exec.Command("notify-send", msg)
			}
		}
		time.Sleep(time.Minute * 15)
	}
}

func GetCollector(urlMap map[string] []string) *Collector{
	c:= &Collector{
		UrlToEndpointsMap: urlMap,
		Client:            &http.Client{},
	}
	var urls []string
	for k, v := range urlMap {
		for _, v1 := range v{
			urls = append(urls, k + v1)
		}
	}
	c.urls = urls
	return c
}
