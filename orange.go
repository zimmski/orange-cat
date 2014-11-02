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

	watcher := NewWatcher(o.filepath)
	watcher.Start()

	mdConverter := NewMdConverter(watcher.GetDataChan(), o.useBasic)

	mdChan := mdConverter.GetMdChan()
	httpServer := NewHTTPServer(portStr, Template(o.filepath, port), mdChan)
	httpServer.Listen()

	open.Run("http://localhost" + portStr)

	<-done

	watcher.Stop()
}
