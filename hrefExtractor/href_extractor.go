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
      if isAnchor {
          fmt.Println("found a link")
      }
    }
  }
}
