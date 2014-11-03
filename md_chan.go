package main

import (
	"github.com/russross/blackfriday"
)

type MdChan struct {
	data    chan *string
	request chan bool
}

func (md *MdChan) MarkdownConverter(rawDataChan chan *string, useBasic bool) {
	var convert func([]byte) []byte
	if useBasic {
		convert = blackfriday.MarkdownBasic
	} else {
		convert = blackfriday.MarkdownCommon
	}

	for {
		select {
		case rawData := <-rawDataChan:
			data := string(convert([]byte(*rawData)))
			md.data <- &data
		default:
		}
	}
}

func NewMdChan(watcherDataChan *DataChan, useBasic bool) *MdChan {
	mdChan := MdChan{make(chan *string), watcherDataChan.request}

	go mdChan.MarkdownConverter(watcherDataChan.data, useBasic)

	return &mdChan
}
