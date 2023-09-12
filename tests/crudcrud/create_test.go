package crudcrud_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/nggrjh/automation-test/steps/crudcrud"
)

var _ = ginkgo.Describe("Create Resource API", func() {
	const url = "https://crudcrud.com/api/1ca2f1c240c94615a0eaecbd4db45a06/unicorns"

	ginkgo.When("", func() {

		ginkgo.It("should create a new resource and return 201 Created", func() {
			requestData := map[string]interface{}{
				"name":   "Sparkle Angel",
				"age":    2,
				"colour": "blue",
			}

			resp, err := crudcrud.Create(url, requestData)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusCreated))

			var responseData map[string]interface{}
			gomega.Expect(json.NewDecoder(resp.Body).Decode(&responseData)).NotTo(gomega.HaveOccurred())

			gomega.Expect(responseData["name"]).To(gomega.Equal("Sparkle Angel"))
			gomega.Expect(responseData["age"]).To(gomega.Equal(float64(2)))
			gomega.Expect(responseData["colour"]).To(gomega.Equal("blue"))
		})

		ginkgo.It("should handle invalid data and return 400 Bad Request", func() {
			// Test with invalid data that should result in a 400 Bad Request
			// Add assertions to validate the error response
		})
	})
})

func TestCreateResource(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Create Resource API Suite")
}
