package main

import (
	"net/http"
)

type Websocket struct {
	markdown <-chan *string
}

func NewWebsocket(markdown <-chan *string) *Websocket {
	return &Websocket{markdown}
}

func (ws *Websocket) Serve(w *http.ResponseWriter, r *http.Request) {
}
