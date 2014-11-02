package main

type MdChan struct {
	data    chan *string
	request chan bool
}

func NewMdChan(watcherDataChan *DataChan, useBasic bool) *MdChan {
	// FIXME: implement Markdown Converter
	dataChan := watcherDataChan.data

	return &MdChan{dataChan, watcherDataChan.request}
}
