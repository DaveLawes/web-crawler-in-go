package hrefExtractor

import (
  "io"
	"golang.org/x/net/html"
  "net/url"
  // "fmt"
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

      if !isValid(href) {
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

func isValid(link string) (success bool) {
  parsed, err := url.Parse(link)
  if err == nil && len(parsed.Hostname()) == 0 && !isPhone(link) {
    if link[0:1] == "/" {
        success = true
    }
  }
  return
}

func isPhone(link string) (isPhone bool) {
  if len(link) > 4 {
    if link[0:4] == "tel:" {
      isPhone = true
    }
  }
  return
}
