package orange_test

import (
	. "."

	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	Describe("#Template()", func() {
		It("should write the result of the template to a writer.", func() {
			w := TestResponseWriter{}

			Template(&w, "temp_file.md", 1234)

			Expect(strings.Contains(w.data, "<title>temp_file.md</title>")).
				To(Equal(true))
			Expect(strings.Contains(w.data, "var conn = new WebSocket(\"ws://localhost:1234/temp_file.md\");")).
				To(Equal(true))
		})

		It("should use a custom CSS if it exists.", func() {
			if CustomCSSPath() == "" {
				// If there's no custom css path, just skip this test.
				return
			}

			// Without a custom CSS
			notExist := os.Rename(CustomCSSPath(), CustomCSSPath()+"_")
			w := TestResponseWriter{}
			Template(&w, "temp_file.md", 1234)
			Expect(strings.Contains(w.data, "<style>")).To(Equal(true))

			// With a custom CSS
			os.Create(CustomCSSPath())
			w = TestResponseWriter{}
			Template(&w, "temp_file.md", 1234)
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
