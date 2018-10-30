package hrefExtractor

import (
  "io"
  "fmt"
	"golang.org/x/net/html"
  "net/url"
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

      href, success := getLinkFromTag(tag)

      if success {
        fmt.Println(href)
      }

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

func isInDomain(link string) bool {
  parsed, error := url.Parse(link)
  if error == nil {
    fmt.Println(parsed.Hostname())
    fmt.Println(error)
  }

  return false
}
