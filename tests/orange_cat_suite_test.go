package orange_cat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOrangeCat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "orange-cat Suite")
}
