package urlPrinter

import (
  "sort"
)

func Print(urlMap map[string][]string, seedUrl string) string {
  output := "-- SEED: " + seedUrl + " --\n\n"
  orderedKeys := order(urlMap, seedUrl)
  for _, key := range orderedKeys {
    output += "- " + key + ":\n"
    for i := 0; i < len(urlMap[key]); i++ {
      output += "    " + urlMap[key][i] + "\n"
    }
    output += "\n"
  }
  output += "-- END --\n"
  return output
}

func order(urlMap map[string][]string, seedUrl string) []string {
  orderedKeys := []string{}
  for k, _ := range urlMap {
      if k != seedUrl {
        orderedKeys = append(orderedKeys, k)
      }
  }
  sort.Strings(orderedKeys)
  orderedKeys = append([]string{seedUrl}, orderedKeys...)
  return orderedKeys
}
