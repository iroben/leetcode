package main

import "testing"

/**
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type TwoSumData struct {
  Nums   []int
  Target int
  Result []int
}

func TestTwoSum(t *testing.T) {
  data := []TwoSumData{
    {Nums: []int{2, 7, 11, 15}, Target: 9, Result: []int{0, 1}},
  }
  for _, d := range data {
    result := TwoSum(d.Nums, d.Target)
    if result == nil || len(result) != 2 {
      t.Error("fail: result nil or len no equal 2")
    }
    if result[0] != d.Result[0] || result[1] != d.Result[1] {
      t.Error("fail")
    }
  }

}
