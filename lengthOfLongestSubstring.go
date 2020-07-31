package main

func LengthOfLongestSubstring(s string) int {
  if s == "" {
    return 0
  }
  matchMap := make(map[byte]int)
  max := 0
  diff := 0
  lastLeft := 0
  l := len(s)
  for i := 0; i < l; i++ {
    left, ok := matchMap[s[i]]
    if ok {
      if left < lastLeft {
        left = lastLeft
      }
      diff = i - left
      lastLeft = left
    } else {
      diff += 1
    }
    if diff > max {
      max = diff
    }
    matchMap[s[i]] = i
  }
  return max
}
