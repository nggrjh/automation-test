package crudcrud

import (
	"net/http"

	"github.com/onsi/gomega"
)

func Read(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	return resp, nil
}
