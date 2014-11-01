package main

import (
	"fmt"
	"time"
)

const (
	DataChanSize = 3
)

func NewOrange(filepath string) *Orange {
	return &Orange{filepath, false}
}

type Orange struct {
	filepath string
	useBasic bool
}

func (o *Orange) UseBasic() {
	o.useBasic = true
}

func (o *Orange) Run(port int) {
	done := make(chan bool)
	data := make(chan *string, DataChanSize)

	watcher := NewWatcher(o.filepath, data)
	watcher.Start()

	temp := time.NewTicker(time.Millisecond * WatcherInterval)
	go func() {
		for {
			<-temp.C

			str := <-data
			fmt.Println("new data!", *str)
		}
	}()

	<-done

	watcher.Stop()
}
