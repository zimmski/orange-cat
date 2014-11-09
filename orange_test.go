package orange

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"strconv"
	"time"
)

var _ = Describe("Orange", func() {
	var (
		port = 6060
	)

	Describe("#NewOrange()", func() {
		It("should return a new Orange object.", func() {
			orange := NewOrange(port)
			Expect(orange).NotTo(BeNil())
		})
	})

	Describe("#orange.UseBasic()", func() {
		It("should set the useBasic property.", func() {
			orange := NewOrange(port)
			orange.UseBasic()
			// No way to check the property
		})
	})

	Describe("#orange.Run()", func() {
		It("should run a orange server listening on the given port.", func() {
			orange := NewOrange(port)
			result := runAndWait(orange, port)

			Expect(result).To(Equal(true))

			orange.Stop()
		})
	})

	Describe("#orange.Stop()", func() {
		var (
			orange *Orange = nil
		)

		BeforeEach(func() {
			orange = NewOrange(port)
			result := runAndWait(orange, port)
			Expect(result).To(Equal(true))
		})

		It("should stop the orange server running.", func() {
			orange.Stop()
			_, err := http.Get("http://localhost:" + strconv.Itoa(port) + "/ping")
			Expect(err).NotTo(BeNil())
		})

		It("should be able to stop the server several times.", func() {
			orange.Stop()
			result := runAndWait(orange, port)
			Expect(result).To(Equal(true))
			orange.Stop()
			_, err := http.Get("http://localhost:" + strconv.Itoa(port) + "/ping")
			Expect(err).NotTo(BeNil())
		})
	})
})

func runAndWait(orange *Orange, port int) bool {
	go func() {
		orange.Run()
	}()

	// Check if we can connect to it
	isListening := false
	ticker := time.NewTicker(time.Millisecond * 500)
	for i := 0; i < 10; i++ {
		<-ticker.C
		resp, err := http.Get("http://localhost:" + strconv.Itoa(port) + "/ping")
		if err == nil && resp.StatusCode == 200 {
			isListening = true
			break
		}
	}
	ticker.Stop()

	return isListening
}
