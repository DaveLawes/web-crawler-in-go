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

  go func() { urlQueue <- seedUrl }()

  for i := 0; i <= len(urlMap); {
    select {
    case url := <- urlQueue:
      go getLinks(url, client, urlMap, chFinished, urlQueue)
    case <- chFinished:
      fmt.Println("chFinished")
      i++
    }
  }

  fmt.Println(urlMap)
  return
}

func getLinks(url string, client HttpClient, urlMap UrlMap, chFinished chan bool, urlQueue chan string) {
  body := getBody.GetBody(client, url)
  links := hrefExtractor.Extract(body)
  for _, url := range links {
    urlQueue <- url
  }
  urlMap[url] = links
  chFinished <- true
}

func addToMap(urlMap UrlMap, links []string) {
  for _, url := range links {
    if _, ok := urlMap[url]; !ok {
      urlMap[url] = []string{}
    }
  }
}
