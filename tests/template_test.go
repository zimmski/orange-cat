package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"os"
	"os/user"
	"path/filepath"
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

		It("uses a custom CSS if it exists.", func() {
			if CustomCSSPath() == "" {
				// If there's no custom css path, just skip this test.
				return
			}

			// Without a custom CSS
			notExist := os.Rename(CustomCSSPath(), CustomCSSPath()+"_")
			template := Template("temp_file.md", 1234)
			w := TestResponseWriter{}
			template(&w)
			Expect(strings.Contains(w.data, "<style>")).To(Equal(true))

			// With a custom CSS
			os.Create(CustomCSSPath())
			template = Template("temp_file.md", 1234)
			w = TestResponseWriter{}
			template(&w)
			customCSS, _ := CustomCSS()
			Expect(strings.Contains(w.data, *customCSS)).To(Equal(true))

			// Reset the original file or remove test file.
			if notExist != nil {
				os.Remove(CustomCSSPath())
			} else {
				os.Rename(CustomCSSPath()+"_", CustomCSSPath())
			}
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

func CustomCSSPath() string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	return filepath.Join(usr.HomeDir, ".orange-cat.css")
}
