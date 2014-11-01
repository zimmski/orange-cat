package main

import (
	"io/ioutil"
	"os"
	"time"
)

type Watcher struct {
	filepath string
	dataChan chan<- *string
	ticker   *time.Ticker
	done     chan bool
}

func NewWatcher(filepath string, dataChan chan<- *string) *Watcher {
	return &Watcher{filepath, dataChan, nil, nil}
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
			default:
				<-w.ticker.C

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
					w.dataChan <- &data
				}
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.done <- true
	w.ticker.Stop()
}
