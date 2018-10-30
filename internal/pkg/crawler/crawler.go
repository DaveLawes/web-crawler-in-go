package crawler

import (
  "github.com/DaveLawes/webcrawler-in-go/internal/pkg/getBody"
  "fmt"
)

func Crawl(seedUrl string) {

  body := getBody.GetBody(seedUrl)
  fmt.Println(body)
}
