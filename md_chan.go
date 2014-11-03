package main

import (
	"github.com/russross/blackfriday"
)

type MdChan struct {
	data    chan *[]byte
	request chan bool
}

func (md *MdChan) MarkdownConverter(rawDataChan chan *[]byte, useBasic bool) {
	var convert func([]byte) []byte
	if useBasic {
		convert = blackfriday.MarkdownBasic
	} else {
		convert = blackfriday.MarkdownCommon
	}

	for {
		select {
		case raw := <-rawDataChan:
			data := convert(*raw)
			md.data <- &data
		default:
		}
	}
}

func NewMdChan(watcherDataChan *DataChan, useBasic bool) *MdChan {
	mdChan := MdChan{make(chan *[]byte), watcherDataChan.request}

	go mdChan.MarkdownConverter(watcherDataChan.raw, useBasic)

	return &mdChan
}
