package kata

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestIsThisATriangle is TestIsThisATriangle
func TestIsThisATriangle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsThisATriangle Suite")
}

var _ = Describe("Test Example", func() {
	It("works for some examples", func() {
		Expect(IsTriangle(5, 1, 2)).To(Equal(false))
		Expect(IsTriangle(1, 2, 5)).To(Equal(false))
		Expect(IsTriangle(2, 5, 1)).To(Equal(false))
		Expect(IsTriangle(4, 2, 3)).To(Equal(true))
		Expect(IsTriangle(5, 1, 5)).To(Equal(true))
		Expect(IsTriangle(2, 2, 2)).To(Equal(true))
  })
})