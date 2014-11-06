package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var _ = Describe("HTTPServer", func() {
	var (
		port       = 6060
		serverAddr = fmt.Sprintf("http://localhost:%d", port)
	)

	Describe("#NewHTTPServer()", func() {
		It("should return a new HTTPServer object.", func() {
			server := NewHTTPServer(port)
			Expect(server).NotTo(BeNil())
		})
	})

	Describe("#Addr()", func() {
		It("should return a addr string.", func() {
			server := NewHTTPServer(port)
			Expect(server.Addr()).To(Equal(fmt.Sprintf(":%d", port)))
		})
	})

	Describe("#httpServer.ListenAndServe()", func() {
		It("should turn on a server.", func() {
			server := NewHTTPServer(port)
			go server.ListenAndServe()

			isListening := false
			ticker := time.NewTicker(time.Millisecond * 500)
			for i := 0; i < 10; i++ {
				<-ticker.C
				resp, err := http.Get(serverAddr + "/ping")
				if err == nil && resp.StatusCode == 200 {
					isListening = true
					break
				}
			}
			ticker.Stop()

			Expect(isListening).To(Equal(true))

			server.Stop()
		})
	})

	Describe("#httpServer.Listen()", func() {
		It("should turn on a server and wait until it's on.", func() {
			server := NewHTTPServer(port)
			server.Listen()

			isListening := false
			resp, err := http.Get(serverAddr + "/ping")
			if err == nil && resp.StatusCode == 200 {
				isListening = true
			}

			Expect(isListening).To(Equal(true))

			server.Stop()
		})

		It("should serve the template page.", func() {
			readme := "README.md"

			server := NewHTTPServer(port)
			server.Listen()

			isListening := false
			resp, err := http.Get(serverAddr + "/" + readme)
			if err == nil && resp.StatusCode == 200 {
				isListening = true
			}

			Expect(isListening).To(Equal(true))

			defer resp.Body.Close()
			content, _ := ioutil.ReadAll(resp.Body)
			w := TestResponseWriter{}
			Template(&w, readme, port)
			Expect(string(content)).To(Equal(w.data))

			server.Stop()
		})
	})

	Describe("#httpServer.Stop()", func() {
		It("should stop the running server.", func() {
			server := NewHTTPServer(port)
			server.Listen()
			server.Stop()

			_, err := http.Get(serverAddr + "/ping")

			Expect(err).NotTo(BeNil())
		})
	})
})
