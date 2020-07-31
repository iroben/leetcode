package main

import "sort"

type NumCount struct {
  Num   int
  Count int
}

func RearrangeBarcodes(barcodes []int) []int {
  numCountMap := make(map[int]int)
  l := len(barcodes)
  for i := 0; i < l; i++ {
    numCountMap[barcodes[i]] += 1
  }
  nl := len(numCountMap)
  nums := make([]NumCount, nl)
  i := 0
  for k, v := range numCountMap {
    nums[i] = NumCount{
      Num:   k,
      Count: v,
    }
    i += 1
  }
  sort.Slice(nums, func(i, j int) bool {
    return nums[i].Count > nums[j].Count
  })
  retVal := make([]int, l)
  k := 0
  for _, v := range nums {
    num, count := v.Num, v.Count
    for j := 0; j < count; j++ {
      retVal[k] = num
      k += 2
      if k >= l {
        k = 1
      }
    }
  }
  return retVal
}
