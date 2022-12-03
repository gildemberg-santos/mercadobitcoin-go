package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configurations struct {
	Moeda             string
	LastPurchaseOrder float64
	Percentage        float64
	PurchaseValue     float64
	Interval          int
	UrlApi            string
}

type Requisitions struct {
	EndPoint string
}

type TickerLoad struct {
	Buy  string `json:"buy"`
	Date int    `json:"date"`
	High string `json:"high"`
	Last string `json:"last"`
	Low  string `json:"low"`
	Open string `json:"open"`
	Pair string `json:"pair"`
	Sell string `json:"sell"`
	Vol  string `json:"vol"`
}

type Ticker struct {
	Buy  float64 `json:"buy"`
	Date int     `json:"date"`
	High float64 `json:"high"`
	Last float64 `json:"last"`
	Low  float64 `json:"low"`
	Open float64 `json:"open"`
	Pair string  `json:"pair"`
	Sell float64 `json:"sell"`
	Vol  float64 `json:"vol"`
}

func (c *Configurations) SetConfigurations() {
	godotenv.Load()
	c.Moeda = os.Getenv("MOEDA")
	c.LastPurchaseOrder, _ = strconv.ParseFloat(os.Getenv("LAST_PURCHASE_ORDER"), 64)
	c.Percentage, _ = strconv.ParseFloat(os.Getenv("PERCENTAGE"), 64)
	c.PurchaseValue, _ = strconv.ParseFloat(os.Getenv("PURCHASE_VALUE"), 64)
	c.Interval, _ = strconv.Atoi(os.Getenv("INTERVAL"))
	c.UrlApi = os.Getenv("URL_API")
}

func (r *Requisitions) GetRequisitions(endpoint string) []byte {
	r.EndPoint = endpoint

	resp, err := http.Get(r.EndPoint)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return body
}

func GetTicker() Ticker {
	var tickers []TickerLoad
	config := Configurations{}
	config.SetConfigurations()

	requisitions := Requisitions{}
	body := requisitions.GetRequisitions(config.UrlApi + "tickers/?symbols=" + config.Moeda)
	err := json.Unmarshal(body, &tickers)
	if err != nil {
		panic(err)
	}

	if len(tickers) == 0 {
		return Ticker{}
	}

	ticker := Ticker{}
	ticker.Buy, _ = strconv.ParseFloat(tickers[0].Buy, 64)
	ticker.Date = tickers[0].Date
	ticker.High, _ = strconv.ParseFloat(tickers[0].High, 64)
	ticker.Last, _ = strconv.ParseFloat(tickers[0].Last, 64)
	ticker.Low, _ = strconv.ParseFloat(tickers[0].Low, 64)
	ticker.Open, _ = strconv.ParseFloat(tickers[0].Open, 64)
	ticker.Pair = tickers[0].Pair
	ticker.Sell, _ = strconv.ParseFloat(tickers[0].Sell, 64)
	ticker.Vol, _ = strconv.ParseFloat(tickers[0].Vol, 64)

	return ticker
}
