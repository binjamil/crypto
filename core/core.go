package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"github.com/m7shapan/njson"
)

const (
	URL     = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"
	API_KEY = "86772962-6647-4a85-a6b2-af41725c1c03"
)

type WebClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CryptoService struct {
	client WebClient
}

type Quote struct {
	Symbol           string
	Price            float32 `njson:"data.*.quote.USD.price"`
	PercentChange1H  float32 `njson:"data.*.quote.USD.percent_change_1h"`
	PercentChange24H float32 `njson:"data.*.quote.USD.percent_change_24h"`
	PercentChange7D  float32 `njson:"data.*.quote.USD.percent_change_7d"`
	PercentChange30D float32 `njson:"data.*.quote.USD.percent_change_30d"`
}

func NewCryptoService(client WebClient) *CryptoService {
	return &CryptoService{client}
}

func (cs *CryptoService) GetQuote(symbol string) (*Quote, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("convert", "USD")
	q.Add("symbol", symbol)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", API_KEY)
	req.URL.RawQuery = q.Encode()

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got error status code %v", res.Status)
	}

	rawJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quote Quote
	quote.Symbol = symbol
	njson.Unmarshal(rawJson, &quote)
	return &quote, nil
}

func (cs *CryptoService) GetQuotes(symbols ...string) map[string]Quote {
	quoteMap := make(map[string]Quote)
	wg := sync.WaitGroup{}

	for _, symbol := range symbols {
		wg.Add(1)
		go func(s string) {
			quote, err := cs.GetQuote(s)
			if err != nil {
				fmt.Printf("Failed to GET %v. Got err %v\n", s, err)
				wg.Done()
				return
			}
			quoteMap[s] = *quote
			wg.Done()
		}(symbol)
	}

	wg.Wait()
	return quoteMap
}
