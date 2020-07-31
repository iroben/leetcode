package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
  data := [][]*ListNode{
    {
      &ListNode{
        Val: 2,
        Next: &ListNode{
          Val: 4,
          Next: &ListNode{
            Val:  3,
            Next: nil,
          },
        },
      }, &ListNode{
      Val: 5,
      Next: &ListNode{
        Val: 6,
        Next: &ListNode{
          Val:  4,
          Next: nil,
        },
      },
    }, &ListNode{
      Val: 7,
      Next: &ListNode{
        Val: 0,
        Next: &ListNode{
          Val:  8,
          Next: nil,
        },
      },
    },


      &ListNode{
        Val:  5,
        Next: nil,
      }, &ListNode{
      Val:  5,
      Next: nil,
    }, &ListNode{
      Val: 0,
      Next: &ListNode{
        Val:  1,
        Next: nil,
      },
    },
    },
  }
  for _, d := range data {
    result := AddTwoNumbers(d[0], d[1])
    if !(result.Val == d[2].Val &&
      result.Next.Val == d[2].Next.Val &&
      result.Next.Next.Val == d[2].Next.Next.Val) {
      t.Errorf("result= %+v,data=%+v\n", result, d)
    }
  }

}
