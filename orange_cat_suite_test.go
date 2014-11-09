package orange

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrangeCat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "orange-cat Suite")
}
