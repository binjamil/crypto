package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/binjamil/crypto/core"
)

func main() {
	cs := core.NewCryptoService(&http.Client{
		Timeout: time.Second * 10,
	})

	quote, err := cs.GetQuote("BTC")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(quote)
	}

	fmt.Println(cs.GetQuotes("ADA", "DOT"))
}
