package orange

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MdConverter", func() {
	Describe("#NewMarkdownConverter()", func() {
		It("should return a new MarkdownConverter object.", func() {
			mdConverter := NewMarkdownConverter()
			Expect(mdConverter).NotTo(BeNil())
		})
	})

	Describe("#mdConverter.UseBasic()", func() {
		It("should set the converter function to the Basic one.", func() {
			mdConverter := NewMarkdownConverter()
			mdConverter.UseBasic()
			// No way to check
		})
	})

	Describe("#mdConverter.Convert()", func() {
		It("should convert a Markdown raw data to HTML format.", func() {
			mdConverter := NewMarkdownConverter()
			raw := []byte("# Hello, `world`")
			expected := []byte("<h1>Hello, <code>world</code></h1>\n")
			Expect(mdConverter.Convert(raw)).To(Equal(expected))
		})

		It("should convert a Basic Markdown raw data.", func() {
			mdConverter := NewMarkdownConverter()
			mdConverter.UseBasic()
			raw := []byte("# Hello, `world`")
			expected := []byte("<h1>Hello, <code>world</code></h1>\n")
			Expect(mdConverter.Convert(raw)).To(Equal(expected))
		})
	})
})
