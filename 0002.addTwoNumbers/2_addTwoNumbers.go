package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	cur := l1
	for {
		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}
		cur.Val += carry
		carry = cur.Val / 10
		cur.Val = cur.Val % 10

		if (carry > 0 || l2 != nil || cur.Next != nil) == false {
			break
		}

		if cur.Next == nil {
			cur.Next = &ListNode{0, nil}
		}

		cur = cur.Next
	}

	return l1
}
