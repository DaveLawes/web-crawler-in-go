package getBody

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

func GetBody(client HttpClient, url string) (httpBody io.ReadCloser) {
  response, err := client.Get(url)
  check(err)
  return response.Body
}
