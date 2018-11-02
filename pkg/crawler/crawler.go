package crawler

import (
  "net/http"
  "sync"
  "web-crawler-in-go/pkg/getBody"
  "web-crawler-in-go/pkg/hrefExtractor"
)

type HttpClient interface {
  Get(string) (*http.Response, error)
}

type UrlMap map[string][]string

func Crawl(seedUrl string, client HttpClient) (urlMap UrlMap) {
  urlMap = make(UrlMap)
  urlQueue := make(chan string)
  chFinished := make(chan bool)

  var waitgroup sync.WaitGroup
  waitgroup.Add(1)


  go func() { urlQueue <- seedUrl }()

  go func() {
    defer waitgroup.Done()
    for current_seed := range urlQueue {
      body := getBody.GetBody(client, current_seed)
      links := hrefExtractor.Extract(body)
      urlMap[current_seed] = links
      chFinished <- true
    }
  }()

  select {
  // case url := <- urlQueue:
  //   fmt.Println("url added to queue")
  //   fmt.Println(url)
  case <- chFinished:
    break
  }

  waitgroup.Wait()
  return
}

func addToMap(urlMap UrlMap, links []string) {
  for _, url := range links {
    if _, ok := urlMap[url]; !ok {
      urlMap[url] = []string{}
    }
  }
}
