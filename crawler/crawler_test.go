package crawler

import (
  "testing"
  "fmt"
  "io/ioutil"
  "bytes"
  "net/http"
)

type MockHttpClient struct {}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
  response := &http.Response {
    Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test"))),
  }
  return response, nil
}

func TestCrawl(t *testing.T) {
  httpClient := &MockHttpClient{}
  returned := Crawl(httpClient, "http://example.com/")

  buf := new(bytes.Buffer)
  buf.ReadFrom(returned)

  response := buf.String()
  expectation := "Test"

  if response != expectation {
    t.Errorf("Return is incorrect, got: %s, want: %s", response, expectation)
  }
}
