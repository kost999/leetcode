package main

import (
	"fmt"
	"strconv"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	carry := 0
//	head := new(ListNode)
//	cur := head
//	for l1 != nil || l2 != nil || carry != 0 {
//		n1, n2 := 0, 0
//		if l1 != nil {
//			n1, l1 = l1.Val, l1.Next
//		}
//		if l2 != nil {
//			n2, l2 = l2.Val, l2.Next
//		}
//		num := n1 + n2 + carry
//		carry = num / 10
//		cur.Next = &ListNode{num % 10, nil}
//		cur = cur.Next
//	}
//	return head.Next
//}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

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

func TestCase1(t *testing.T) {
	/*
		Input: l1 = [2,4,3], l2 = [5,6,4]
		Output: [7,0,8]
		Explanation: 342 + 465 = 807.
	*/
	//l111 := ListNode{2, nil}
	//l11 := ListNode{4, &l111}
	//l1 := ListNode{3, &l11}
	//l211 := ListNode{5, nil}
	//l21 := ListNode{6, &l211}
	//l2 := ListNode{4, &l21}
	//fmt.Print(printListNode(addTwoNumbers(&l1, &l2)))
	//
	//c2l1 := ListNode{0, nil}
	//c2l2 := ListNode{0, nil}
	//fmt.Print(printListNode(addTwoNumbers(&c2l1, &c2l2)))

	c3l1111111 := ListNode{9, nil}
	c3l111111 := ListNode{9, &c3l1111111}
	c3l11111 := ListNode{9, &c3l111111}
	c3l1111 := ListNode{9, &c3l11111}
	c3l111 := ListNode{9, &c3l1111}
	c3l11 := ListNode{9, &c3l111}
	c3l1 := ListNode{9, &c3l11}

	c3l2111 := ListNode{9, nil}
	c3l211 := ListNode{9, &c3l2111}
	c3l21 := ListNode{9, &c3l211}
	c3l2 := ListNode{9, &c3l21}

	fmt.Print(printListNode(addTwoNumbers(&c3l1, &c3l2)))
}

func printListNode(listNode *ListNode) (ret string) {
	if listNode != nil {
		for {
			ret += ", " + strconv.Itoa(listNode.Val)
			if listNode.Next == nil {
				break
			}
			listNode = listNode.Next
		}
	}

	return
}
