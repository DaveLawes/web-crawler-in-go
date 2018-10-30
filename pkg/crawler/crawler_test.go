package crawler

import (
  "testing"
  "net/http"
  "io/ioutil"
  "fmt"
  "bytes"
  "reflect"
)

type MockNewGetBody struct {}

type MockNewHrefExtractor struct {}

type MockHttpClient struct {}

// type UrlMap map[string][]string

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
  fmt.Println(url)
  response := &http.Response {
    Body: ioutil.NopCloser(bytes.NewBuffer([]byte(`<a class="c-header__button" href="/download">Sign up</a>`))),
  }
  return response, nil
}

func TestCrawler_Crawl(t *testing.T) {
  httpClient := &MockHttpClient{}
  result := Crawl("http://example.com", httpClient)

  expectation := make(UrlMap)
  links := []string { "/download" }
  expectation["http://example.com"] = links

  if reflect.DeepEqual(result, expectation) == false {
    t.Errorf("Result does not match expectation")
  }
}
