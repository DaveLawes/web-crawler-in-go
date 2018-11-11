package urlPrinter

import (
  "testing"
  "io/ioutil"
  "reflect"
)



func TestUrlPrinter_Print(t *testing.T) {
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

func TestUrlPrinter_order(t *testing.T) {
  urlMap := make(map[string][]string)
  seed := "http://example.com"
  urlMap["/test3"] = []string{}
  urlMap["http://example.com"] = []string{ "/test1", "/test2"}
  urlMap["/test2"] = []string{}
  urlMap["/test1"] = []string{ "/test3" }

  expectation := []string{ "http://example.com", "/test1", "/test2", "/test3" }

  result := order(urlMap, seed)

  if reflect.DeepEqual(result, expectation) == false {
    t.Errorf("alphabeticalOrder incorrect, got: %s, want: %s", result, expectation)
  }

}
