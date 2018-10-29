package hrefExtractor

import (
  "testing"
  "io/ioutil"
  "bytes"
  "strings"
  "fmt"
	"golang.org/x/net/html"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func TestExtractTotal(t *testing.T) {
  buffer, err := ioutil.ReadFile("./test_data/monzo_home_html.txt")
  check(err)

  test_data := bytes.NewReader(buffer)

  result := Extract(test_data)
  expectation := []string{"/", "/download", "/-play-store-redirect",
    "/features/apple-pay", "/features/travel", "/community", "/transparency",
    "blog/how-money-works", "/tone-of-voice", "/faq", "/legal/terms-and-conditions",
    "legal/fscs-information", "/legal/privacy-policy", "/legal/cookie-policy"}

  if len(result) != len(expectation) {
    t.Errorf("Result length was incorrect, got: %d, want: %d", len(result), len(expectation))
  }

}

func TestGetLinkFromTag(t *testing.T) {
  string_data := `<a class="c-header__button" href="/download">Sign up</a>`
  test_data := strings.NewReader(string_data)

  tokenizer := html.NewTokenizer(test_data)
  token := tokenizer.Next()
  tag := tokenizer.Token()

  href_attr := tag.Attr[1]

  fmt.Println(token)
  fmt.Println(href_attr.Val)

  result, _ := GetLinkFromTag(tag)
  fmt.Println(result)

  if result != "/download" {
    t.Errorf("Href is incorrect, got: %s, want: %s", result, href_attr.Val)
  }

}
