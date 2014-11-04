package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"net/http"
	"time"
)

var _ = Describe("HTTPServer", func() {
	var (
		port     = 6060
		portStr  = ":6060"
		template = Template("temp_file.md", port)
		mdChan   *MdChan
	)

	BeforeEach(func() {
		dataChan := NewDataChan()
		mdChan = NewMdChan(dataChan, false)
	})

	Describe("#NewHTTPServer()", func() {
		It("should return a new HTTPServer object.", func() {
			server := NewHTTPServer(portStr, template, mdChan)
			Expect(server).NotTo(BeNil())
		})
	})

	Describe("#httpServer.ListenAndServe()", func() {
		It("should turn on a server.", func() {
			server := NewHTTPServer(portStr, template, mdChan)
			go server.ListenAndServe()

			isListening := false
			ticker := time.NewTicker(time.Millisecond * 500)
			for i := 0; i < 10; i++ {
				<-ticker.C
				resp, err := http.Get("http://localhost" + portStr + "/ping")
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
			server := NewHTTPServer(portStr, template, mdChan)
			server.Listen()

			isListening := false
			resp, err := http.Get("http://localhost" + portStr + "/ping")
			if err == nil && resp.StatusCode == 200 {
				isListening = true
			}

			Expect(isListening).To(Equal(true))

			server.Stop()
		})

		It("should serve the template page.", func() {
			server := NewHTTPServer(portStr, template, mdChan)
			server.Listen()

			isListening := false
			resp, err := http.Get("http://localhost" + portStr)
			if err == nil && resp.StatusCode == 200 {
				isListening = true
			}

			Expect(isListening).To(Equal(true))

			defer resp.Body.Close()
			content, _ := ioutil.ReadAll(resp.Body)
			w := TestResponseWriter{}
			template(&w)
			Expect(string(content)).To(Equal(w.data))

			server.Stop()
		})
	})

	Describe("#httpServer.Stop()", func() {
		It("should stop the running server.", func() {
			server := NewHTTPServer(portStr, template, mdChan)
			server.Listen()
			server.Stop()

			_, err := http.Get("http://localhost" + portStr + "/ping")

			Expect(err).NotTo(BeNil())
		})
	})
})
