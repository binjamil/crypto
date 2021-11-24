package core

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/m7shapan/njson"
)

type WebClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CryptoService struct {
	client WebClient
}

type Quote struct {
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
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("convert", "USD")
	q.Add("symbol", symbol)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "86772962-6647-4a85-a6b2-af41725c1c03")
	req.URL.RawQuery = q.Encode()

	res, err := cs.client.Do(req)
	if err != nil {
		return nil, err
	}

	rawJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quote Quote
	njson.Unmarshal(rawJson, &quote)
	return &quote, nil
}
