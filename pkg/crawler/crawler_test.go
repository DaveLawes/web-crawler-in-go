package crawler

import (
  "testing"
  "net/http"
  "io/ioutil"
  // "fmt"
  "bytes"
)

type MockNewGetBody struct {}

type MockNewHrefExtractor struct {}

type MockHttpClient struct {}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
  response := &http.Response {
    Body: ioutil.NopCloser(bytes.NewBuffer([]byte(`<a class="c-header__button" href="/download">Sign up</a>`))),
  }
  return response, nil
}

func TestCrawler_Crawl(t *testing.T) {
  httpClient := &MockHttpClient{}
  result := Crawl("http://example.com", httpClient)

  expectation := "http://example.com:\n/download\n"

  if result != expectation {
    t.Errorf("Result does not match expectation, got: %s, want: %s", result, expectation)
  }
}
