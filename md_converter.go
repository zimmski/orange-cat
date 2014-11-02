package main

type MdChan struct {
	data    chan *string
	request chan bool
}

type MdConverter struct {
	mdChan *MdChan
}

func NewMdConverter(watcherDataChan *DataChan, useBasic bool) *MdConverter {
	// FIXME: implement Markdown Converter
	dataChan := watcherDataChan.data

	mdChan := MdChan{dataChan, watcherDataChan.request}
	return &MdConverter{&mdChan}
}

func (md *MdConverter) GetMdChan() *MdChan {
	return md.mdChan
}
