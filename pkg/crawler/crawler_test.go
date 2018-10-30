package crawler

import (
  "testing"
  "io"
  "io/ioutil"
  "fmt"
  "bytes"
)

type MockNewGetBody struct {}

func (m *MockNewGetBody) GetBody(client HttpClient, url string) (httpBody io.ReadCloser) {
  body := ioutil.NopCloser(bytes.NewBuffer([]byte("Test")))
  return body
}

func TestCrawler_Crawl(t *testing.T) {
  mockGetBody := &MockNewGetBody{}
  body := Crawl(mockGetBody, "http://example.com")
  fmt.Println(body)
}
