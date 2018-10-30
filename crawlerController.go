package main

import (
  "net/http"
  "io"
)

type HttpClient interface {
   Get(string) (*http.Response, error)
}

func main() {
  seedUrl := "https://monzo.com"
  client := &http.Client{}




}

func crawl() {

}

// href extractor returns an array of url strings
// struct of key value pairs
// key = url string
// value = array received from hrefExtractor
// use keys to determine if site has already been crawled

// initiate vars
// start crawl
// when http body is returned --> extract
// when all urls have been crawled --> print

// create http client and pass in as argument to crawler.crawl

// client := &http.Client{}
