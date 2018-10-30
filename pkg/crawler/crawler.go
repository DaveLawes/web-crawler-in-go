package crawler

import (
  "net/http"
  "fmt"
  "web-crawler-in-go/pkg/getBody"
  "web-crawler-in-go/pkg/hrefExtractor"
)

type HttpClient interface {
  Get(string) (*http.Response, error)
}

type UrlMap map[string][]string

func Crawl(seedUrl string, client HttpClient) (UrlMap) {
  urlMap := make(UrlMap)
  urlQueue := make(chan string)
  chFinished := make(chan bool)

  go func() { urlQueue <- seedUrl }()

  go func() {
    for current_seed := range urlQueue {
      body := getBody.GetBody(client, current_seed)
      links := hrefExtractor.Extract(body)
      urlMap[current_seed] = links
      chFinished <- true
    }
  }()

  select {
  case <- chFinished:
    return urlMap
  }
}

func addToMap(urlMap UrlMap, links []string) {
  for _, url := range links {
    if _, ok := urlMap[url]; !ok {
      urlMap[url] = []string{}
    }
  }
}
