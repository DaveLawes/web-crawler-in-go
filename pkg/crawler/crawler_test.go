package crawler

import (
  "testing"
  "net/http"
  "io/ioutil"
  "fmt"
  "bytes"
  "reflect"
)

type MockHttpClient struct {}

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
  expectation["/download"] = links

  if reflect.DeepEqual(result, expectation) == false {
    t.Errorf("Expected urlMap does not match result. Expected: %v, got: %v", expectation, result)
  }
}

func TestCrawler_getAbsUrl(t *testing.T) {
  seedUrl := "http://example.com"
  url := "/download"
  expectation := "http://example.com/download"
  result := getAbsUrl(seedUrl, url)

  if expectation != result {
    t.Errorf("Incorrect absolute url created. Expected: %s, got: %s", expectation, result)
  }
}

func TestCrawler_getRelUrl(t *testing.T) {
  parentUrl := "/download"
  url := "current"

  expectation := "/download/current"
  result := getRelUrl(url, parentUrl)

  if expectation != result {
    t.Errorf("Incorrect relative url created. Expected: %s, got: %s", expectation, result)
  }
}
