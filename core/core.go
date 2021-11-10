package core

import (
	"net/http"
	"net/url"
	"time"
)

func Fetch() (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	// q.Add("start", "1")
	// q.Add("limit", "5000")
	q.Add("convert", "USD")
	q.Add("symbol", "ADA")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "86772962-6647-4a85-a6b2-af41725c1c03")
	req.URL.RawQuery = q.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
