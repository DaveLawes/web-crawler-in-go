package hrefExtractor

import (
  "testing"
  "io/ioutil"
  // "fmt"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func TestExtract(t *testing.T) {
  test_data, err := ioutil.ReadFile("./test_data/monzo_home_html.txt")
  check(err)

  result := Extract(test_data)
  expectation := [14]string{"/", "/download", "/-play-store-redirect",
    "/features/apple-pay", "/features/travel", "/community", "/transparency",
    "blog/how-money-works", "/tone-of-voice", "/faq", "/legal/terms-and-conditions",
    "legal/fscs-information", "/legal/privacy-policy", "/legal/cookie-policy"}

  if result != expectation {
    t.Errorf("Result was incorrect, got: %s, want: %s", result, expectation)
  }


}
