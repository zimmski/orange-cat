package orange

import (
	"github.com/russross/blackfriday"
)

// Global Converter
var MdConverter = NewMarkdownConverter()

type MarkdownConverter struct {
	convert func([]byte) []byte
}

func NewMarkdownConverter() *MarkdownConverter {
	return &MarkdownConverter{blackfriday.MarkdownCommon}
}

func (md *MarkdownConverter) UseBasic() {
	md.convert = blackfriday.MarkdownBasic
}

func (md *MarkdownConverter) Convert(raw []byte) []byte {
	return md.convert(raw)
}
