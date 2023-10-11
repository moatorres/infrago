package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFileServerHandler(t *testing.T) {
	// setup test server
	ts := httptest.NewServer(FileServerHandler())
	defer ts.Close()

	// send a GET request to the base url
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// check status code
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 HTTP Status Code, received %d", res.StatusCode)
	}

	// check content
	content, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// read the index.html file
	expectedContent, err := os.ReadFile("./static/index.html")
	if err != nil {
		t.Fatal(err)
	}

	// compare with the server response
	if string(content) != string(expectedContent) {
		t.Fatalf("Expected content %q, received %q", expectedContent, content)
	}
}
