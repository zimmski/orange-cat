package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

var _ = Describe("MdChan", func() {
	var (
		dataChan *DataChan
	)

	BeforeEach(func() {
		dataChan = NewDataChan()
	})

	Describe("#NewMdChan()", func() {
		It("should return a new MdChan object.", func() {
			mdChan := NewMdChan(dataChan, false)
			Expect(mdChan).NotTo(BeNil())
			mdChan.Stop()
		})

		It("should automatically process a Markdown raw data.", func() {
			mdChan := NewMdChan(dataChan, false)
			raw := []byte("# Hello, `world`")
			dataChan.Raw <- &raw
			expected := "<h1>Hello, <code>world</code></h1>\n"
			Expect(string(*<-mdChan.Data)).To(Equal(expected))
			mdChan.Stop()
		})
	})

	Describe("#mdChan.MarkdownConverter()", func() {
		It("should process a Markdown raw data.", func() {
			mdChan := NewMdChan(dataChan, false)
			rawDataChan := make(chan *[]byte)

			go mdChan.MarkdownConverter(rawDataChan, false)

			raw := []byte("# Hello, `world`")
			rawDataChan <- &raw
			expected := "<h1>Hello, <code>world</code></h1>\n"
			Expect(string(*<-mdChan.Data)).To(Equal(expected))

			mdChan.Stop()
		})

		It("should process a Basic Markdown raw data.", func() {
			mdChan := NewMdChan(dataChan, true)
			rawDataChan := make(chan *[]byte)

			go mdChan.MarkdownConverter(rawDataChan, true)

			raw := []byte("# Hello, `world`")
			rawDataChan <- &raw
			expected := "<h1>Hello, <code>world</code></h1>\n"
			Expect(string(*<-mdChan.Data)).To(Equal(expected))

			mdChan.Stop()
		})
	})

	Describe("#mdChan.Stop()", func() {
		It("should stop the Markdown process.", func() {
			mdChan := NewMdChan(dataChan, false)
			raw := []byte("# Hello, `world`")

			dataChan.Raw <- &raw
			expected := "<h1>Hello, <code>world</code></h1>\n"
			Expect(string(*getOne(mdChan.Data))).To(Equal(expected))

			mdChan.Stop()

			dataChan.Raw <- &raw
			Expect(getOne(mdChan.Data)).To(BeNil())
		})
	})
})

func getOne(c chan *[]byte) *[]byte {
	<-time.After(time.Millisecond * 500) // Wait enough time
	var result *[]byte
	select {
	case result = <-c:
	default:
		result = nil
	}
	return result
}
