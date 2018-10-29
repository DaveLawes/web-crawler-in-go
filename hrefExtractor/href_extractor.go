package hrefExtractor

import (
  "io"
  "fmt"
	"golang.org/x/net/html"
)

func Extract(responseBody io.Reader) []string {
  result := []string{}
  tokenizer := html.NewTokenizer(responseBody)

  for {
  	token := tokenizer.Next()
  	switch {
    case token == html.ErrorToken:
      return result
    case token == html.StartTagToken:
      tag := tokenizer.Token()
      isAnchor := tag.Data == "a"

      if !isAnchor {
        continue
      }

      href, ok := GetLinkFromTag(tag)

      if ok {
        fmt.Println(href)
      }

    }
  }
}

func GetLinkFromTag(tag html.Token) (href string, success bool) {
  for _, a := range tag.Attr {
      if a.Key == "href" {
        href = a.Val
        success = true
      }
  }
  fmt.Println(href)
  fmt.Println(success)
  return
}
