package main

import "testing"

type MockHttpClient struct {}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
  response := &http.Response {
    Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test"))),
  }
  return response, nil
}

func TestCrawler(t *testing.T) {

  // check seed is added to channel

}

func TestCrawler_crawl(t *testing.T) {
  
}
