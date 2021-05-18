package net_test

import (
	"fmt"
	"io/ioutil"
	"learngo/net"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestHttpRequest(t *testing.T) {
	log.Println("start test http request")
	s := net.NewMyServer()

	go s.StartHttpServer()

	log.Printf("server: %v\n", s)
	// tt := getTestcases(s.Port, s.Timeout)
	// tc := tt[0]

	// change port to see different resp err
	link := fmt.Sprintf("http://localhost:%d00", s.Port)
	// link := tc.link

	// change timeout to see different resp err
	// timeout := (s.HandleTime + 1) * time.Second
	timeout := 3 * time.Second

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatalf("failed to create request url %s\n%#v", link, err)
	}

	c := &http.Client{Timeout: timeout}
	resp, err := c.Do(req)
	// resp, err := http.Get(link)
	log.Printf("resp: %#v\n", resp)
	if err != nil {
		if err, ok := err.(*url.Error); ok {
			log.Printf("is timeout: %v\n", err.Timeout())
		}
		log.Fatalf("failed to request url %s\n%#v", link, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(link, "bytes", len(body))
	log.Println("message", string(body))
}

type testCase struct {
	link    string
	timeout time.Duration
	resp    http.Response
	respErr error
}

func getTestcases(port int, timeout time.Duration) []testCase {
	return []testCase{
		{
			link:    fmt.Sprintf("http://localhost:%d", port),
			timeout: (timeout + 1) * time.Second,
			resp:    http.Response{StatusCode: 200},
			respErr: nil,
		},
	}
}
