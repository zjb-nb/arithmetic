package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/swap-nodes-in-pairs/description/
[1,2,3,4]
[2,1,4,3]
扩展：25 K个一组翻转链表
*/
func SwapPairs(head *ListNode) *ListNode {
	//TODO
	panic("")
}

/*
[1,2,3,4]
[2,1,4,3]
1.先两辆截取21 和43  不断更待 cur和next
2.再分别旋转
*/
func swapPairs(head *ListNode) *ListNode {
	prev := new(ListNode)
	res := prev
	for head != nil {
		i := 0
		start := head
		for i < 2 && head != nil {
			head = head.Next
			if head == nil {
				prev.Next = start
				return res.Next
			}
			i++
		}
		next := head.Next
		head.Next = nil
		prev.Next = swap(start)
		prev = start
		head = next
	}
	return res.Next
}

func swapPairs2(head *ListNode) *ListNode {
	//TODO
	prev := &ListNode{}
	prev.Next = head
	res := prev

	for head != nil && head.Next != nil {
		next := head.Next
		temp := next.Next

		prev.Next = next
		head.Next.Next = head
		head.Next = temp

		prev = head
		head = temp
	}
	return res.Next
}

// 递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	head.Next = swapPairs1(res.Next)
	res.Next = head
	return res
}

func swap(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := swap(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func TestSwap(t *testing.T) {
	tests := []struct {
		name    string
		data    *ListNode
		wantRes []int
	}{
		{
			name:    "1",
			data:    MakeList([]int{1, 2, 3, 4}),
			wantRes: []int{2, 1, 4, 3},
		},
		{
			name:    "2",
			data:    MakeList([]int{1}),
			wantRes: []int{1},
		},
		{
			name:    "1",
			data:    MakeList([]int{}),
			wantRes: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SwapPairs(tt.data).Array()
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
