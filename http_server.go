package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	ListeningTestInterval = 500
	MaxListeningTestCount = 10
)

type HTTPServer struct {
	port     string
	template func(http.ResponseWriter)
	ws       *Websocket
	listener net.Listener
}

func NewHTTPServer(port string, template func(http.ResponseWriter), mdChan *MdChan) *HTTPServer {
	return &HTTPServer{port, template, NewWebsocket(mdChan), nil}
}

func (s *HTTPServer) ListenAndServe() {
	var err error
	server := &http.Server{
		Addr:           s.port,
		Handler:        s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.listener, err = net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}

	server.Serve(s.listener)
}

func (s *HTTPServer) Listen() {
	go s.ListenAndServe()

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

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" {
		w.Write([]byte("pong"))
	} else if r.URL.Path == "/ws" {
		s.ws.Serve(w, r)
	} else {
		s.template(w)
	}
}

func (s *HTTPServer) Stop() {
	s.listener.Close()
}
