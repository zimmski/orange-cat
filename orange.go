package main

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

	httpServer := NewHttpServer(port, Template(o.filepath, port), data)
	httpServer.Listen()

	<-done

	watcher.Stop()
}
