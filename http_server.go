package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type HttpServer struct {
	dataChan <-chan *string
	server   *http.Server
}

func NewHttpServer(dataChan <-chan *string) *HttpServer {
	return &HttpServer{dataChan, nil}
}

func (s *HttpServer) Listen(port int) {
	portStr := ":" + strconv.Itoa(port)

	s.server = &http.Server{
		Addr:           portStr,
		Handler:        s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.server.ListenAndServe()
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
		// FIXME
		w.Write([]byte("Hello, world!"))
	}
}
