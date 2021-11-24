package main

import (
	"log"
	"net/http"
	"time"

	"github.com/binjamil/crypto/core"
)

func main() {
	cs := core.NewCryptoService(&http.Client{
		Timeout: time.Second * 10,
	})

	_, err := cs.GetQuote("DOT")
	if err != nil {
		log.Fatal(err)
	}
}
