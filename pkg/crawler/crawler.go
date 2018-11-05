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

func Crawl(seedUrl string, client HttpClient) (urlMap UrlMap) {
  urlMap = make(UrlMap)
  urlQueue := make(chan string)
  chFinished := make(chan bool)

  pChFinished := &chFinished

  go func() { urlQueue <- seedUrl }()

  // go func() {
  //   // for current_seed := range urlQueue {
  //   //   body := getBody.GetBody(client, current_seed)
  //   //   links := hrefExtractor.Extract(body)
  //   //   urlMap[current_seed] = links
  //   //   chFinished <- true
  //   // }
  //   getLinks(urlQueue, client, urlMap, chFinished)
  // }()
  for i := 0; i < 1; i++ {
    select {
    case url := <- urlQueue:
      fmt.Println("url added to queue")
      fmt.Println(url)
      getLinks(url, client, urlMap, pChFinished)
    case <- chFinished:
      fmt.Println("chFinished")
      i++
      // break
    }
  }
  return
}

func getLinks(url string, client HttpClient, urlMap UrlMap, pChFinished *chan bool) {
  fmt.Println("getLinks")
  fmt.Println(url)
  // for current_seed := range urlQueue {
    body := getBody.GetBody(client, url)
    fmt.Println(body)
    links := hrefExtractor.Extract(body)
    fmt.Println(links)
    urlMap[url] = links
    *pChFinished <- true

  // }
}

func addToMap(urlMap UrlMap, links []string) {
  for _, url := range links {
    if _, ok := urlMap[url]; !ok {
      urlMap[url] = []string{}
    }
  }
}
