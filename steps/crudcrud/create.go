package crudcrud

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/onsi/gomega"
)

func Create(url string, requestData map[string]any) (*http.Response, error) {
	requestBody, err := json.Marshal(requestData)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	resp, err := http.Post(url, "application/json", bytes.NewReader(requestBody))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	return resp, nil
}
