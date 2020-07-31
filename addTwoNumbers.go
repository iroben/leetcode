package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
  Val  int
  Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
  n1, n2 := l1, l2
  retVal := new(ListNode)
  current := retVal
  carry := 0
  for n1 != nil || n2 != nil {
    v1, v2 := 0, 0
    if n1 != nil {
      v1 = n1.Val
      n1 = n1.Next
    }
    if n2 != nil {
      v2 = n2.Val
      n2 = n2.Next
    }
    sum := carry + v1 + v2
    carry = sum / 10
    current.Next = new(ListNode)
    current = current.Next
    current.Val = sum % 10
  }
  if carry > 0 {
    current.Next = new(ListNode)
    current.Next.Val = carry
  }
  return retVal.Next
}
