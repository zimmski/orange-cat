package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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

	fmt.Println("Listening", portStr, "...")
}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ws" {
		// FIXME
		http.Error(w, "Not yet implemented...", 404)
	} else {
		s.template(&w)
	}
}
