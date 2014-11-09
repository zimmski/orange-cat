package orange

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Websocket", func() {
	Describe("#NewWebsocket()", func() {
		It("should return a new Websocket object.", func() {
			sock := NewWebsocket("README.md")
			Expect(sock).NotTo(BeNil())
		})
	})
})
