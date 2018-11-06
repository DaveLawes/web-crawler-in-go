package urlPrinter

import (
  "fmt"
)

type UrlMap map[string][]string

func Print(urlMap UrlMap, seedUrl string) string {
  output := "-- SEED: " + seedUrl + " --\n"
  for key, value := range urlMap {
    output += key + " :\n"
    fmt.Println(value)
  }
  return output
}
