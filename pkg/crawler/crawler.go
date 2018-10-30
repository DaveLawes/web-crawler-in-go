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
  urlMap := make(UrlMap)

  body := getBody.GetBody(client, seedUrl)
  links := hrefExtractor.Extract(body)
  fmt.Println(links)
  urlMap[seedUrl] = links

  return urlMap
}

// map:
// key       value
// parent: [children],
// parent: [children]

// put seed in queue
// iterate over queue:
//    invoke GetBody and Extract
// know when finished
// close queue
