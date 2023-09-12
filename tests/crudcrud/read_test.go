package crudcrud_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/nggrjh/automation-test/steps/crudcrud"
	"github.com/nggrjh/automation-test/steps/util"
)

var _ = ginkgo.Describe("Read Resource API", func() {
	const url = "https://crudcrud.com/api/1ca2f1c240c94615a0eaecbd4db45a06/unicorns"

	ginkgo.When("does not require created data", func() {
		ginkgo.It("should handle the case when a resource is not found and return 404 Not Found", func() {
			resp, err := http.Get(url + "/" + util.NewHex(24))
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusNotFound))
		})
	})

	ginkgo.When("requires created data", func() {
		var createdId string

		ginkgo.BeforeEach(func() {
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

			createdId = responseData["_id"].(string)
			gomega.Expect(createdId).NotTo(gomega.BeEmpty())
		})

		ginkgo.It("should retrieve an existing resource and return 200 OK", func() {
			resp, err := http.Get(url + "/" + createdId)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			gomega.Expect(resp.StatusCode).To(gomega.Equal(http.StatusOK))

			// var responseData map[string]interface{}
			// gomega.Expect(json.NewDecoder(resp.Body).Decode(&responseData)).NotTo(gomega.HaveOccurred())

			// gomega.Expect(responseData["name"]).To(gomega.Equal("Sparkle Angel"))
			// gomega.Expect(responseData["age"]).To(gomega.Equal(float64(2)))
			// gomega.Expect(responseData["colour"]).To(gomega.Equal("blue"))
		})
	})

})

func TestReadResource(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Read Resource API Suite")
}
