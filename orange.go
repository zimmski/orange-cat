package main

import (
	"github.com/skratchdot/open-golang/open"
)

const (
	MarkdownChanSize = 3
)

func NewOrange(filepath string) *Orange {
	return &Orange{filepath, false, make(chan chan<- bool)}
}

type Orange struct {
	filepath string
	useBasic bool
	stop     chan chan<- bool
}

func (o *Orange) UseBasic() {
	o.useBasic = true
}

func (o *Orange) Run(port int) {
	watcher := NewWatcher(o.filepath)
	watcher.Start()

	mdChan := NewMdChan(watcher.GetDataChan(), o.useBasic)

	httpServer := NewHTTPServer(o.filepath, port, mdChan)
	httpServer.Listen()

	open.Run("http://localhost" + httpServer.PortStr())

	done := <-o.stop

	httpServer.Stop()
	mdChan.Stop()
	watcher.Stop()

	done <- true
}

func (o *Orange) Stop() {
	done := make(chan bool)
	o.stop <- done
	<-done
}
