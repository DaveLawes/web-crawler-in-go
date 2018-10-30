package crawler

import (
  "net/http"
  "io"
  "fmt"
  "web-crawler-in-go/pkg/getBody"
  "web-crawler-in-go/pkg/hrefExtractor"
)

type HttpClient interface {
  Get(string) (*http.Response, error)
}

type NewGetBody interface {
  GetBody(client HttpClient, url string) (httpBody io.ReadCloser)
}

func Crawl(seedUrl string, client HttpClient) string {
  foundUrls := []string {}

  body := getBody.GetBody(client, seedUrl)
  links := hrefExtractor.Extract(body)
  fmt.Println(links)
  foundUrls = append(foundUrls, links...)

  output := ""
  for i := 0; i < len(foundUrls); i++ {
    output += (foundUrls[i] + "\n")
  }

  return fmt.Sprintf("%s:\n%s", seedUrl, output)
}

// put seed in queue
// iterate over queue:
//    invoke GetBody and Extract
// know when finished
// close queue
