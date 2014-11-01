package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	ListeningTestInterval = 500
	MaxListeningTestCount = 10
)

type HttpServer struct {
	port     string
	template func(*http.ResponseWriter)
	ws       *Websocket
}

func NewHttpServer(port string, template func(*http.ResponseWriter), dataChan <-chan *string) *HttpServer {
	return &HttpServer{port, template, NewWebsocket(dataChan)}
}

func (s *HttpServer) Listen() {
	server := &http.Server{
		Addr:           s.port,
		Handler:        s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	isListening := make(chan bool)
	go func() {
		result := false
		ticker := time.NewTicker(time.Millisecond * ListeningTestInterval)
		for i := 0; i < MaxListeningTestCount; i++ {
			<-ticker.C
			resp, err := http.Get("http://localhost" + s.port + "/ping")
			if err == nil && resp.StatusCode == 200 {
				result = true
				break
			}
		}
		ticker.Stop()
		isListening <- result
	}()

	if <-isListening {
		fmt.Println("Listening", s.port, "...")
	} else {
		panic("Can't connect to server")
	}
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" {
		w.Write([]byte("pong"))
	} else if r.URL.Path == "/ws" {
		s.ws.Serve(&w, r)
	} else {
		s.template(&w)
	}
}
