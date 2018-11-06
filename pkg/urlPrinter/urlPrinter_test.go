package urlPrinter

import (
  "testing"
  "io/ioutil"
)



func TestUrlPrinter_Print(t *testing.T) {
  // desired output is a .txt file with the pretty printed map
  urlMap := make(map[string][]string)
  urlMap["http://example.com"] = []string{ "/test1", "/test2"}
  urlMap["/test1"] = []string{ "/test3" }
  urlMap["/test2"] = []string{}
  urlMap["/test3"] = []string{}

  seedUrl := "http://example.com"
  byteVals, _ := ioutil.ReadFile("../../test/test_data/test_print.txt")
  expectation := string(byteVals)

  result := Print(urlMap, seedUrl)

  if result != expectation {
    t.Errorf("Returned string doesn't match expectation, got: %s, want: %s", result, expectation)
  }
}
