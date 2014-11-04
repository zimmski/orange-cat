package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Websocket", func() {
	var (
		mdChan *MdChan
	)

	BeforeEach(func() {
		dataChan := NewDataChan()
		mdChan = NewMdChan(dataChan, false)
	})

	Describe("#NewWebsocket()", func() {
		It("should return a new Websocket object.", func() {
			sock := NewWebsocket(mdChan)
			Expect(sock).NotTo(BeNil())
		})
	})
})
