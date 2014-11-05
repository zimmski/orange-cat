package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"reflect"
	"strings"
)

var _ = Describe("Template", func() {
	Describe("#Template()", func() {
		It("should return a template function.", func() {
			template := Template("temp_file.md", 1234)
			Expect(reflect.TypeOf(template).String()).
				To(Equal("func(http.ResponseWriter)"))
		})

		It("can be used as a template function.", func() {
			template := Template("temp_file.md", 1234)
			w := TestResponseWriter{}
			template(&w)

			Expect(strings.Contains(w.data, "<title>temp_file.md</title>")).
				To(Equal(true))
			Expect(strings.Contains(w.data, "var conn = new WebSocket(\"ws://localhost:1234/ws\");")).
				To(Equal(true))
		})
	})
})

type TestResponseWriter struct {
	data string
}

func (w *TestResponseWriter) Header() http.Header { return nil }
func (w *TestResponseWriter) Write(data []byte) (int, error) {
	w.data = string(data)
	return len(data), nil
}
func (w *TestResponseWriter) WriteHeader(int) {}
