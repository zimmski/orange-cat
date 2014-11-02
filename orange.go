package main

import (
	"github.com/skratchdot/open-golang/open"
	"strconv"
)

const (
	MarkdownChanSize = 3
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
	markdown := make(chan *string, MarkdownChanSize)

	watcher := NewWatcher(o.filepath, markdown)
	watcher.Start()

	httpServer := NewHTTPServer(portStr, Template(o.filepath, port), markdown)
	httpServer.Listen()

	open.Run("http://localhost" + portStr)

	<-done

	watcher.Stop()
}
