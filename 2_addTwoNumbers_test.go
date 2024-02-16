package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestCase1(t *testing.T) {
	/*
		input
		l1 = [2,4,3]
		l2 = [5,6,4]
		output
		[7,0,8]
	*/
	assert.Equal(t,
		&ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
		addTwoNumbers(
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
		),
	)
}

func TestCase2(t *testing.T) {
	/*
		input
		l1 = [0]
		l2 = [0]
		output
		[0]
	*/
	assert.Equal(t,
		&ListNode{
			Val:  0,
			Next: nil,
		},
		addTwoNumbers(
			&ListNode{
				Val:  0,
				Next: nil,
			},
			&ListNode{
				Val:  0,
				Next: nil,
			},
		),
	)
}

func TestCase3(t *testing.T) {
	/*
		input
		l1 = [9,9,9,9,9,9,9]
		l2 = [9,9,9,9]
		output
		[8,9,9,9,0,0,0,1]
	*/
	assert.Equal(t,
		&ListNode{
			Val: 8,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val:  1,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
		addTwoNumbers(
			&ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val:  9,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			&ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
		),
	)
}
