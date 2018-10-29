package hrefExtractor

import (
  "testing"
  "io/ioutil"
  "bytes"
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
