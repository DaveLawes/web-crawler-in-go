package crawler

import (
  "net/http"
  "io"
)

type HttpClient interface {
   Get(string) (*http.Response, error)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func Crawl(client HttpClient, url string) (httpBody io.ReadCloser) {
  response, error := client.Get(url)
  check(error)

  return response.Body
}
