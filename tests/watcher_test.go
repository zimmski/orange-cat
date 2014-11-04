package orange_cat_test

import (
	. "../"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
	"time"
)

var _ = Describe("Watcher", func() {
	var (
		testFilepath = "./test_file"
	)

	BeforeEach(func() {
		// Create a test file
		err := ioutil.WriteFile(testFilepath, []byte("Hello, world"), 0644)
		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		// Remove a test file
		err := os.Remove(testFilepath)
		if err != nil {
			panic(err)
		}
	})

	Describe("#NewWatcher()", func() {
		It("should return a new Watcher object.", func() {
			watcher := NewWatcher(testFilepath)
			Expect(watcher).NotTo(BeNil())
		})
	})

	Describe("#watcher.GetDataChan()", func() {
		It("should return a DataChan object.", func() {
			watcher := NewWatcher(testFilepath)
			Expect(watcher.GetDataChan()).NotTo(BeNil())
		})
	})

	Describe("#watcher.Start()", func() {
		It("should start watching the file modification.", func() {
			watcher := NewWatcher(testFilepath)
			watcher.Start()

			c := watcher.GetDataChan()
			Expect(*<-c.Raw).To(Equal([]byte("Hello, world")))

			watcher.Stop()
		})

		It("should convey a new data if the file is modified.", func() {
			watcher := NewWatcher(testFilepath)
			watcher.Start()

			c := watcher.GetDataChan()
			Expect(*<-c.Raw).To(Equal([]byte("Hello, world")))

			modify(testFilepath, "Hi, there")
			Expect(*<-c.Raw).To(Equal([]byte("Hi, there")))

			watcher.Stop()
		})

		It("should convey a data again when it's requested.", func() {
			watcher := NewWatcher(testFilepath)
			watcher.Start()

			c := watcher.GetDataChan()
			Expect(*<-c.Raw).To(Equal([]byte("Hello, world")))

			c.Request <- true
			Expect(*<-c.Raw).To(Equal([]byte("Hello, world")))

			watcher.Stop()
		})
	})

	Describe("#watcher.Stop()", func() {
		It("should stop watching.", func() {
			watcher := NewWatcher(testFilepath)
			watcher.Start()

			c := watcher.GetDataChan()
			Expect(*<-c.Raw).To(Equal([]byte("Hello, world")))

			watcher.Stop()

			modify(testFilepath, "Hi, there")
			var data *[]byte
			select {
			case data = <-c.Raw:
			default:
				data = nil
			}
			Expect(data).To(BeNil())
		})
	})
})

func modify(testFilepath string, content string) {
	<-time.After(time.Millisecond * 1100) // Wait until ModTime changes
	err := ioutil.WriteFile(testFilepath, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
