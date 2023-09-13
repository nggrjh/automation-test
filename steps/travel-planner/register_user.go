package travelplanner

import (
	"bytes"
	"net/http"

	"github.com/nggrjh/automation-test/utils"
	"github.com/onsi/gomega"
)

func RegisterUser(url, query string, variables map[string]any) (*http.Response, error) {
	body := map[string]any{
		"query":     query,
		"variables": variables,
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(utils.ToByte(body)))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())

	return resp, nil
}
