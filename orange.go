package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
)

const (
	MarkdownChanSize = 3
)

func NewOrange(port int) *Orange {
	return &Orange{port, nil, make(chan bool)}
}

type Orange struct {
	port       int
	httpServer *HTTPServer
	stop       chan bool
}

func (o *Orange) UseBasic() {
	MdConverter.UseBasic()
}

func (o *Orange) Run(files ...string) {
	o.httpServer = NewHTTPServer(o.port)
	o.httpServer.Listen()

	for _, file := range files {
		addr := fmt.Sprintf("http://localhost:%d/%s", o.port, file)
		open.Run(addr)
	}

	<-o.stop
}

func (o *Orange) Stop() {
	o.httpServer.Stop()
	o.stop <- true
}
