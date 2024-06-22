package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/reverse-linked-list-ii/
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: -1, Next: head}
	leftNode := dummy
	for i := 0; i < left-1; i++ {
		if leftNode.Next == nil {
			return head
		}
		leftNode = leftNode.Next
	}
	tailNode, headNode := leftNode.Next, leftNode.Next
	for i := 0; i < right-left; i++ {
		if tailNode.Next == nil {
			return head
		}
		tailNode = tailNode.Next
	}

	rightNode := tailNode.Next
	tailNode.Next = nil
	leftNode.Next = reverseList4(leftNode.Next)
	headNode.Next = rightNode
	return dummy.Next
}

// 头插法思路，只遍历一遍
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	leftNode := dummy
	for i := 0; i < left-1; i++ {
		leftNode = leftNode.Next
	}

	curr := leftNode.Next
	for i := 0; i < right-left; i++ {
		//依靠 curr.Next推进循环
		next := curr.Next
		curr.Next = next.Next
		next.Next = leftNode.Next
		leftNode.Next = next
	}
	return dummy.Next
}

func ReverseBetweenOnce(head *ListNode, step int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//需要保存两个节点的信息即 leftNode 和 rightNode
	leftNode, rightNode := head, head
	for i := 0; i < step-1; i++ {
		if rightNode.Next == nil {
			return head
		}
		rightNode = rightNode.Next
	}
	rightNode.Next = nil
	ReverseList(leftNode)

	return rightNode
}

func TestReverseBetween1(t *testing.T) {
	head := MakeList([]int{1, 2, 3, 4, 5})
	newhead := reverseBetween2(head, 2, 4)
	t.Log(newhead.Array())
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		left    int
		right   int
		wantRes []int
	}{
		{
			name:    "[1,2,3,4,5]",
			data:    []int{1, 2, 3, 4, 5},
			left:    2,
			right:   4,
			wantRes: []int{1, 4, 3, 2, 5},
		},
		{
			name:    "[5]",
			data:    []int{5},
			left:    1,
			right:   1,
			wantRes: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := MakeList(tt.data)
			reslist := reverseBetween(head, tt.left, tt.right)
			res := reslist.Array()
			assert.Equal(t, tt.wantRes, res)
		})
	}

}
