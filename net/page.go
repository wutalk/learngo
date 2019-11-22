package net

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadPage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url, "bytes",len(body))
}
