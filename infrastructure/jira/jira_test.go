package jira

import (
	"net/http"
	"net/http/httptest"
)

func setup() (*http.ServeMux, *httptest.Server) {
	// Test server
	testMux := http.NewServeMux()
	testServer := httptest.NewServer(testMux)
	return testMux, testServer
}
