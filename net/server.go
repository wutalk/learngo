package net

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("req header: %#v\n", req.Header)
	msg := []byte("Hello, web!\n")
	_, err := w.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("reqest handled")
}

func StartHttpServer() {
	http.HandleFunc("/", viewHandler)
	hostUrl := "localhost:8080"
	fmt.Println("Server is listening at", hostUrl)
	err := http.ListenAndServe(hostUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
}
