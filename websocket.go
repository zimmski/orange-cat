package main

import (
	"net/http"
)

type Websocket struct {
	dataChan <-chan *string
}

func NewWebsocket(dataChan <-chan *string) *Websocket {
	return &Websocket{dataChan}
}

func (ws *Websocket) Serve(w *http.ResponseWriter, r *http.Request) {
}
