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

func Crawl(seedUrl string, client HttpClient) UrlMap {
  fmt.Println("Crawl")
  urlMap := make(UrlMap)
  urlQueue := make(chan string)
  go func() { urlQueue <- seedUrl }()

  go func() {
    fmt.Println("inside go func")
    for current_seed := range urlQueue {
      fmt.Println("inside for loop")
      body := getBody.GetBody(client, current_seed)
      links := hrefExtractor.Extract(body)
      fmt.Println(links)
      urlMap[current_seed] = links
    }
  }()

  return urlMap
}

// put seed in queue
// iterate over queue:
//    invoke GetBody and Extract
// know when finished
// close queue

// first step:
// add seed to queue
// use url in queue inside a function
// update from within func
// return same result
