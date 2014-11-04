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
	Raw     chan *[]byte
	Request chan bool
}

type Watcher struct {
	filepath string
	dataChan *DataChan
	ticker   *time.Ticker
	stop     chan bool
}

func NewDataChan() *DataChan {
	return &DataChan{make(chan *[]byte, DataChanSize), make(chan bool)}
}

func NewWatcher(filepath string) *Watcher {
	return &Watcher{filepath, NewDataChan(), nil, nil}
}

func (w *Watcher) Start() {
	go func() {
		w.ticker = time.NewTicker(time.Millisecond * WatcherInterval)
		defer w.ticker.Stop()
		w.stop = make(chan bool)
		var currentTimestamp int64 = 0
		for {
			select {
			case <-w.stop:
				return
			case <-w.ticker.C:
				var reload bool = false
				select {
				case <-w.dataChan.Request:
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

					w.dataChan.Raw <- &raw
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
