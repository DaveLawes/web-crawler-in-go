package hrefExtractor

import (
  "io"
	"golang.org/x/net/html"
  "net/url"
)

func Extract(responseBody io.Reader) []string {
  result := []string {}
  tokenizer := html.NewTokenizer(responseBody)

  for {
  	token := tokenizer.Next()
  	switch {
    case token == html.ErrorToken:
      return result
    case token == html.StartTagToken:
      tag := tokenizer.Token()

      if tag.Data != "a" {
        continue
      }

      href, success := getLinkFromTag(tag)

      if !success {
        continue
      }

      if !isInDomain(href) {
        continue
      }
      
      result = append(result, href)
    }
  }
}

func getLinkFromTag(tag html.Token) (href string, success bool) {
  for _, a := range tag.Attr {
      if a.Key == "href" {
        href = a.Val
        success = true
      }
  }
  return
}

func isInDomain(link string) (success bool) {
  parsed, error := url.Parse(link)
  if error == nil && len(parsed.Hostname()) == 0 {
    success = true
  }
  return
}