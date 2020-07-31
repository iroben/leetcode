package main

import "testing"

func TestRearrangeBarcodes(t *testing.T) {
  data := [][][]int{
    {
      []int{
        1, 1, 1, 2, 2, 2,
      },
      {
        1, 2, 1, 2, 1, 2,
      },
    },
    {
      []int{
        1, 1, 1, 1, 2, 2, 3, 3,
      },
      {
        1, 2, 1, 2, 1, 3, 1, 3,
      },
    },
  }
  for _, d := range data {
    result := RearrangeBarcodes(d[0])
    for i, v := range result {
      if d[1][i] != v {
        t.Errorf("error: v = %v, d[1] = %v ", result, d[1])
      }
    }
  }
}
