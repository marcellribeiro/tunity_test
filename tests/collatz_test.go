package collatz_test

import (
	"encoding/json"
	"fmt"
	"github.com/marcellribeiro/tunity_test/pkg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCollatz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Collatz Suite")
}

var _ = Describe("Collatz conjecture test", func() {
	Describe("Given number", func() {
		Context("Number is 10", func() {
			It("Result is wrong with number 10", func() {
				expectedResult := "[{\"Step\":1,\"Result\":10},{\"Step\":2,\"Result\":5},{\"Step\":3,\"Result\":16},{\"Step\":4,\"Result\":8},{\"Step\":5,\"Result\":4},{\"Step\":6,\"Result\":2},{\"Step\":7,\"Result\":1}]"
				calcResult, _ := pkg.GetCollatz(10)
				marshaledJson, err := json.Marshal(calcResult)
				if err != nil {
					fmt.Println(err)
				}

				Expect(string(marshaledJson)).To(Equal(expectedResult))
			})
		})
	})

	Describe("Given number", func() {
		Context("Number is 5", func() {
			It("Result is wrong with number 5", func() {
				expectedResult := "[{\"Step\":1,\"Result\":5},{\"Step\":2,\"Result\":16},{\"Step\":3,\"Result\":8},{\"Step\":4,\"Result\":4},{\"Step\":5,\"Result\":2},{\"Step\":6,\"Result\":1}]"
				calcResult, _ := pkg.GetCollatz(5)
				marshaledJson, err := json.Marshal(calcResult)
				if err != nil {
					fmt.Println(err)
				}

				Expect(string(marshaledJson)).To(Equal(expectedResult))
			})
		})
	})

	Describe("Given negative number", func() {
		Context("Number is -20", func() {
			It("Result is wrong with number -20", func() {
				expectedResult := "given number must be greater than 0"
				_, err := pkg.GetCollatz(-20)

				Expect(err).NotTo(Equal(nil))
				Expect(err.Error()).To(Equal(expectedResult))
			})
		})
	})

})
