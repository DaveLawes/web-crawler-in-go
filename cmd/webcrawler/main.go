package main

import (
  "net/http"
  "fmt"
  "os"
  "github.com/DaveLawes/web-crawler-in-go/pkg/crawler"
  "github.com/DaveLawes/web-crawler-in-go/pkg/urlPrinter"
)

func main() {
  fmt.Println("crawling in progress...")
  seedUrl := "https://monzo.com"
  client := &http.Client{}
  result := crawler.Crawl(seedUrl, client)
  stringToWrite := urlPrinter.Print(result, seedUrl)
  text_file, _ := os.Create("Site map.txt")
  defer text_file.Close()
  text_file.WriteString(stringToWrite)
  fmt.Println("Finished! See 'Site map.txt' for result")
}
