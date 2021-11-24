package core_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/binjamil/crypto/core"
)

type MockClient struct {
}

func (mock *MockClient) Do(req *http.Request) (*http.Response, error) {
	json := `{"status":{"timestamp":"2021-11-23T06:33:14.251Z","error_code":0,"error_message":null,"elapsed":23,"credit_count":1,"notice":null},"data":{"DOT":{"id":6636,"name":"Polkadot","symbol":"DOT","slug":"polkadot-new","num_market_pairs":258,"date_added":"2020-08-19T00:00:00.000Z","tags":["substrate","polkadot","binance-chain","binance-smart-chain","polkadot-ecosystem","three-arrows-capital-portfolio","polychain-capital-portfolio","arrington-xrp-capital-portfolio","blockchain-capital-portfolio","boostvc-portfolio","cms-holdings-portfolio","coinfund-portfolio","fabric-ventures-portfolio","fenbushi-capital-portfolio","hashkey-capital-portfolio","kinetic-capital","1confirmation-portfolio","placeholder-ventures-portfolio","pantera-capital-portfolio","exnetwork-capital-portfolio"],"max_supply":null,"circulating_supply":987579314.957085,"total_supply":1103303471.382273,"is_active":1,"platform":null,"cmc_rank":8,"is_fiat":0,"last_updated":"2021-11-23T06:32:07.000Z","quote":{"USD":{"price":39.193630093771105,"volume_24h":1157624395.2213328,"volume_change_24h":-18.7546,"percent_change_1h":-0.41328151,"percent_change_24h":-1.77462663,"percent_change_7d":-6.45389681,"percent_change_30d":-9.9465318,"percent_change_60d":20.22409577,"percent_change_90d":49.52094879,"market_cap":38706818358.68786,"market_cap_dominance":1.5209,"fully_diluted_market_cap":43242468138.53,"last_updated":"2021-11-23T06:32:07.000Z"}}}}}
	&{39.19363 -0.4132815 -1.7746266 -6.453897 -9.946532}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestGetQuote(t *testing.T) {
	cs := core.NewCryptoService(&MockClient{})
	quote, err := cs.GetQuote("DOT")
	if err != nil {
		t.Fatalf("Got error %v", err)
	}

	if quote.Price == 0.0 {
		t.Fatalf("Got zero valued quote struct: %v", quote)
	}
}
