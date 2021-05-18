package net

import (
	"log"
	"net"
	"net/http"
	"time"
)

type MyServer struct {
	Port       int
	HandleTime time.Duration
	Listener   net.Listener
}

func NewMyServer() *MyServer {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	host := listener.Addr().(*net.TCPAddr).IP
	port := listener.Addr().(*net.TCPAddr).Port

	// fmt.Printf("Server is listening at: %s:%d\n", host, port)
	log.Printf("Server is listening at: %s:%d\n", host, port)
	return &MyServer{
		Port:       port,
		HandleTime: 2,
		Listener:   listener,
	}
}

func (s *MyServer) StartHttpServer() {
	// hostUrl := "localhost:8090"

	// err = http.ListenAndServe(hostUrl, nil)
	http.HandleFunc("/", s.viewHandler)
	err := http.Serve(s.Listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MyServer) viewHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("req header: %#v\n", req.Header)
	msg := []byte("Hello, web!\n")
	// make it slow
	log.Println("complex computing...")
	time.Sleep(s.HandleTime * time.Second)

	_, err := w.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("reqest handled")
}
