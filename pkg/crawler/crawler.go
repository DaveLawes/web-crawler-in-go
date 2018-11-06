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
  urlCrawled := make(chan bool)
  crawlComplete := false
  i := 0

  go func() { urlQueue <- seedUrl }()

  for crawlComplete == false {
    select {
    case url := <- urlQueue:
      go getLinks(url, client, urlMap, urlCrawled, urlQueue, seedUrl)
    case <- urlCrawled:
      fmt.Println("urlCrawled")
      i++
      fmt.Println(i)
      fmt.Println(len(urlQueue))
      fmt.Println(urlMap)
      fmt.Println(len(urlMap))
      if i == len(urlMap) {
        fmt.Println("all urls crawled")
        crawlComplete = true
      }
    }
  }

  return
}

func getLinks(url string, client HttpClient, urlMap UrlMap, urlCrawled chan bool, urlQueue chan string, seed string) {
  absUrl := getAbsUrl(seed, url)
  fmt.Println(absUrl)
  body := getBody.GetBody(client, absUrl)
  links := hrefExtractor.Extract(body)
  addToMap(url, urlMap, links, urlQueue)
  urlCrawled <- true
}

func addToMap(url string, urlMap UrlMap, links []string, urlQueue chan string) {
  fmt.Println("addToMap")
  relUrls := getRelUrls(url, links)
  urlMap[url] = relUrls
  for _, url := range relUrls {
    if _, ok := urlMap[url]; !ok {
      urlMap[url] = []string{}
      fmt.Println("url added to queue: ", url)
      urlQueue <- url
    }
  }
}

func getRelUrls(url string, links []string) (relUrls []string) {
  for i := 0; i < len(links); i++ {
    if links[i][0:1] == "/" {
      relUrls = append(relUrls, links[i])
    } else {
      relUrls = append(relUrls, url + "/" + links[i])
    }
  }
  return
}

func getAbsUrl(seedUrl string, url string) (absUrl string) {
  if url != seedUrl {
    absUrl = seedUrl + url
  } else {
    absUrl = seedUrl
  }
  return
}
