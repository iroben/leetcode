package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
  data := map[string]int{
    "abcabcbb": 3,
    "bbbbb":    1,
    "pwwkew":   3,
    " ":        1,
    "abba":     2,
    "tmmzuxt":  5,
    "wobgrovw": 6,
  }
  for k, v := range data {
    v1 := LengthOfLongestSubstring(k)
    if v1 != v {
      t.Errorf("error: %s v = %d, result = %d", k, v, v1)
    }
  }
}
