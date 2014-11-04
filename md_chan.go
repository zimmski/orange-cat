package main

import (
	"github.com/russross/blackfriday"
)

type MdChan struct {
	data    chan *[]byte
	request chan bool
	stop    chan bool
}

func (md *MdChan) MarkdownConverter(rawDataChan chan *[]byte, useBasic bool) {
	md.stop = make(chan bool)

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
		case <-md.stop:
			return
		default:
		}
	}
}

func (md *MdChan) Stop() {
	md.stop <- true
}

func NewMdChan(watcherDataChan *DataChan, useBasic bool) *MdChan {
	mdChan := MdChan{make(chan *[]byte), watcherDataChan.request, nil}

	go mdChan.MarkdownConverter(watcherDataChan.raw, useBasic)

	return &mdChan
}
