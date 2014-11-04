package main

import (
	"github.com/skratchdot/open-golang/open"
	"strconv"
)

const (
	MarkdownChanSize = 3
)

func NewOrange(filepath string) *Orange {
	return &Orange{filepath, false, nil}
}

type Orange struct {
	filepath string
	useBasic bool
	stop     chan bool
}

func (o *Orange) UseBasic() {
	o.useBasic = true
}

func (o *Orange) Run(port int) {
	portStr := ":" + strconv.Itoa(port)

	o.stop = make(chan bool)

	watcher := NewWatcher(o.filepath)
	watcher.Start()
	defer watcher.Stop()

	mdChan := NewMdChan(watcher.GetDataChan(), o.useBasic)
	defer mdChan.Stop()

	httpServer := NewHTTPServer(portStr, Template(o.filepath, port), mdChan)
	httpServer.Listen()
	defer httpServer.Stop()

	open.Run("http://localhost" + portStr)

	<-o.stop
}

func (o *Orange) Stop() {
	o.stop <- true
}
