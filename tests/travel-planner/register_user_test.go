package travelplanner_test

import (
	"encoding/json"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	travelplanner "github.com/nggrjh/automation-test/steps/travel-planner"
	"github.com/nggrjh/automation-test/types"
)

var _ = ginkgo.Describe(types.SpecRegisterUserQuery, func() {
	const url = "http://localhost:8080/graphql"
	const query = `
mutation RegisterUser($email: String!, $password: String!) { 
	registerUser(email: $email, password: $password) {
		email 
	} 
}`

	ginkgo.When("empty email", func() {
		ginkgo.It("should return error", func() {
			request := map[string]any{
				"email":    "",
				"password": "",
			}

			resp, err := travelplanner.RegisterUser(url, query, request)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			var response types.GraphResponse
			gomega.Expect(json.NewDecoder(resp.Body).Decode(&response)).NotTo(gomega.HaveOccurred())

			gomega.Expect(response.Errors).To(gomega.HaveLen(1))
			gomega.Expect(response.Errors[0].Message).To(gomega.Equal("invalid email"))
		})
	})

	ginkgo.When("not empty email, empty password", func() {
		ginkgo.It("should return error", func() {
			request := map[string]any{
				"email":    "test@gmail.com",
				"password": "",
			}

			resp, err := travelplanner.RegisterUser(url, query, request)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			var response types.GraphResponse
			gomega.Expect(json.NewDecoder(resp.Body).Decode(&response)).NotTo(gomega.HaveOccurred())

			gomega.Expect(response.Errors).To(gomega.HaveLen(1))
			gomega.Expect(response.Errors[0].Message).To(gomega.Equal("invalid password"))
		})
	})

	ginkgo.When("not empty email, not empty password", func() {
		ginkgo.It("should return ok", func() {
			request := map[string]any{
				"email":    "test@gmail.com",
				"password": "test",
			}

			resp, err := travelplanner.RegisterUser(url, query, request)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
			defer resp.Body.Close()

			var response types.GraphResponse
			gomega.Expect(json.NewDecoder(resp.Body).Decode(&response)).NotTo(gomega.HaveOccurred())

			gomega.Expect(response.Errors).To(gomega.HaveLen(0))
		})
	})
})

func TestRegisterUser(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, types.SpecRegisterUserQuery)
}
