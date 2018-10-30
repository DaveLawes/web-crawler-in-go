package main

import (
  "net/http"
  "web-crawler-in-go/pkg/crawler"
)

func main() {
  seedUrl := "https://monzo.com"
  client := &http.Client{}
  crawler.Crawl(seedUrl, client)
}
