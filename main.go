package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/binjamil/crypto/core"
)

func main() {
	res, err := core.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Status)
	respBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(respBody))
}
