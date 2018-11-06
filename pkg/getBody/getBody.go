package getBody

import (
  "net/http"
  "io"
  "fmt"
)

type HttpClient interface {
   Get(string) (*http.Response, error)
}

func check(e error) {
  if e != nil {
    fmt.Println("error!")
    return
  }
}

func GetBody(client HttpClient, url string) (httpBody io.ReadCloser) {
  response, err := client.Get(url)
  check(err)
  // defer response.Body.Close()
  return response.Body
}
