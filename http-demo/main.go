package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"net/http"
	"net/url"
)

func main() {
	fmt.Println("start test http request")
	link := "http://localhost:37910"

	c := &http.Client{Timeout: 1 * time.Second}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatalf("failed to create request url %s\n%#v", link, err)
	}
	resp, err := c.Do(req)
	// resp, err := http.Get(link)
	if err != nil {
		if err, ok := err.(*url.Error); ok {
			fmt.Printf("is timeout: %v\n", err.Timeout())
		}
		log.Fatalf("failed to request url %s\n%#v", link, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(link, "bytes", len(body))
	fmt.Println("message", string(body))
}
