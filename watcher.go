package main

import (
	"io/ioutil"
	"os"
	"time"
)

const (
	WatcherInterval = 500
)

type Watcher struct {
	filepath string
	markdown chan<- *string
	ticker   *time.Ticker
	done     chan bool
}

func NewWatcher(filepath string, markdown chan<- *string) *Watcher {
	return &Watcher{filepath, markdown, nil, nil}
}

func (w *Watcher) Start() {
	w.ticker = time.NewTicker(time.Millisecond * WatcherInterval)
	w.done = make(chan bool)
	go func() {
		var currentTimestamp int64 = 0
		for {
			select {
			case <-w.done:
				return
			case <-w.ticker.C:
				info, err := os.Stat(w.filepath)
				if err != nil {
					continue
				}

				timestamp := info.ModTime().Unix()
				if currentTimestamp < timestamp {
					currentTimestamp = timestamp

					raw, err := ioutil.ReadFile(w.filepath)
					if err != nil {
						continue
					}

					data := string(raw)
					w.markdown <- &data
				}
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.done <- true
	w.ticker.Stop()
}
