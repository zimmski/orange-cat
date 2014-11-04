package main

import (
	"io/ioutil"
	"os"
	"time"
)

const (
	WatcherInterval = 500
	DataChanSize    = 10
)

type DataChan struct {
	raw     chan *[]byte
	request chan bool
}

type Watcher struct {
	filepath string
	dataChan *DataChan
	ticker   *time.Ticker
	stop     chan bool
}

func NewWatcher(filepath string) *Watcher {
	dataChan := DataChan{make(chan *[]byte, DataChanSize), make(chan bool)}
	return &Watcher{filepath, &dataChan, nil, nil}
}

func (w *Watcher) Start() {
	w.ticker = time.NewTicker(time.Millisecond * WatcherInterval)
	defer w.ticker.Stop()
	w.stop = make(chan bool)
	go func() {
		var currentTimestamp int64 = 0
		for {
			select {
			case <-w.stop:
				return
			case <-w.ticker.C:
				var reload bool = false
				select {
				case <-w.dataChan.request:
					reload = true
				default:
				}

				info, err := os.Stat(w.filepath)
				if err != nil {
					continue
				}

				timestamp := info.ModTime().Unix()
				if currentTimestamp < timestamp || reload {
					currentTimestamp = timestamp

					raw, err := ioutil.ReadFile(w.filepath)
					if err != nil {
						continue
					}

					w.dataChan.raw <- &raw
				}
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.stop <- true
}

func (w *Watcher) GetDataChan() *DataChan {
	return w.dataChan
}
