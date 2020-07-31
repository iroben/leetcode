package main
/**
1. 两数之和
 */
func TwoSum(nums []int, target int) []int {
  cache := make(map[int]int)
  for i, num := range nums {
    if k, ok := cache[target-num]; ok {
      return []int{k, i}
    } else {
      cache[num] = i
    }
  }
  return nil
}
