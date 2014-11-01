package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	ListeningTestInterval = 500
	MaxListeningTestCount = 10
)

type HttpServer struct {
	port     int
	template func(*http.ResponseWriter)
	dataChan <-chan *string
}

func NewHttpServer(port int, template func(*http.ResponseWriter), dataChan <-chan *string) *HttpServer {
	return &HttpServer{port, template, dataChan}
}

func (s *HttpServer) Listen() {
	portStr := ":" + strconv.Itoa(s.port)

	server := &http.Server{
		Addr:           portStr,
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
			resp, err := http.Get("http://localhost" + portStr + "/ping")
			if err == nil && resp.StatusCode == 200 {
				result = true
				break
			}
		}
		ticker.Stop()
		isListening <- result
	}()

	if <-isListening {
		fmt.Println("Listening", portStr, "...")
	} else {
		panic("Can't connect to server")
	}
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" {
		w.Write([]byte("pong"))
	} else if r.URL.Path == "/ws" {
		// FIXME
		http.Error(w, "Not yet implemented...", 404)
	} else {
		s.template(&w)
	}
}
