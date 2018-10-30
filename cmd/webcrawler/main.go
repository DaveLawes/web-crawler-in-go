package main

import (
  "net/http"
  "io"
  "getBody"
)

type HttpClient interface {
   Get(string) (*http.Response, error)
}

func main() {
  seedUrl := "https://monzo.com"
  client := &http.Client{}
}
