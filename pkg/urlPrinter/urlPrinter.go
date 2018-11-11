package urlPrinter

import (
  "fmt"
  "sort"
)

type UrlMap map[string][]string

func Print(urlMap UrlMap, seedUrl string) string {
  output := "-- SEED: " + seedUrl + " --\n"
  orderedKeys := alphabeticalOrder(urlMap)
  for _, key := range orderedKeys {
    fmt.Println(key)
    output += key + " :\n"
    for i := 0; i < len(urlMap[key]); i++ {
      output += "   " + urlMap[key][i] + "\n"
    }
    output += "\n"
  }
  return output
}

func alphabeticalOrder(urlMap UrlMap) []string {
  orderedKeys := make([]string, len(urlMap))
  i := 0
  for k, _ := range urlMap {
      orderedKeys[i] = k
      i++
  }
  sort.Strings(orderedKeys)
  return orderedKeys
}
