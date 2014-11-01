package main

import (
	"github.com/skratchdot/open-golang/open"
	"strconv"
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
	portStr := ":" + strconv.Itoa(port)

	done := make(chan bool)
	data := make(chan *string, DataChanSize)

	watcher := NewWatcher(o.filepath, data)
	watcher.Start()

	httpServer := NewHttpServer(portStr, Template(o.filepath, port), data)
	httpServer.Listen()

	open.Run("http://localhost" + portStr)

	<-done

	watcher.Stop()
}
