package main

import (
	"github.com/russross/blackfriday"
)

type MdChan struct {
	Data    chan *[]byte
	Request chan bool
	stop    chan bool
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
			md.Data <- &data
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
	mdChan := MdChan{make(chan *[]byte), watcherDataChan.Request, make(chan bool)}

	go mdChan.MarkdownConverter(watcherDataChan.Raw, useBasic)

	return &mdChan
}
