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
      fmt.Println("end of go func")
      fmt.Println(urlMap)
      chFinished <- true
    }
  }()

  select {
  case <- chFinished:
    return urlMap
  }
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
