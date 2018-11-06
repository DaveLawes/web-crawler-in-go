package main

import (
  "net/http"
  "fmt"
  "github.com/DaveLawes/web-crawler-in-go/pkg/crawler"
)

func main() {
  seedUrl := "https://monzo.com"
  client := &http.Client{}
  result := crawler.Crawl(seedUrl, client)
  fmt.Println(result)
}
