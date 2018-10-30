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

func TestExtract(t *testing.T) {
  buffer, err := ioutil.ReadFile("./test_data/monzo_home_html.txt")
  check(err)
  test_data := bytes.NewReader(buffer)

  result := Extract(test_data)
  expectation := []string {"/", "/download", "/-play-store-redirect",
    "/features/apple-pay", "/features/travel", "/community", "/transparency",
    "blog/how-money-works", "/tone-of-voice", "/faq", "/legal/terms-and-conditions",
    "legal/fscs-information", "/legal/privacy-policy", "/legal/cookie-policy"}

  if len(result) != len(expectation) {
    t.Errorf("Result length was incorrect, got: %d, want: %d", len(result), len(expectation))
  }

}

func TestExtract_getLinkFromTag(t *testing.T) {
  string_data := `<a class="c-header__button" href="/download">Sign up</a>`
  test_data := strings.NewReader(string_data)
  tokenizer := html.NewTokenizer(test_data)
  token := tokenizer.Next()
  fmt.Println(token)
  tag := tokenizer.Token()

  result, _ := getLinkFromTag(tag)
  expectation := "/download"

  if result != expectation {
    t.Errorf("Href is incorrect, got: %s, want: %s", result, expectation)
  }

}

func TestExtract_isInDomain_internal(t *testing.T) {
  internal_link := "/download"
  internal_result := isInDomain(internal_link)
  internal_expectation := true

  if internal_result != internal_expectation {
    t.Errorf("Result is incorrect, got: %v, want %v", internal_result, internal_expectation)
  }
}

func TestExtract_isInDomain_external(t *testing.T) {
  external_link := "https://itunes.apple.com/gb/app/mondo/id1052238659"
  external_result := isInDomain(external_link)
  external_expectation := false

  if external_result != external_expectation {
    t.Errorf("Result is incorrect, got: %v, want %v", external_result, external_expectation)
  }
}
