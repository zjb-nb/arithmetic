package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/merge-two-sorted-lists/description/
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
	list1.Next = MergeTwoLists(list2, list1.Next)
	return list1
}

/*
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l := new(ListNode)
	res := l
	for list1 != nil || list2 != nil {
		l.Next = new(ListNode)
		l = l.Next
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			l.Val = list2.Val
			list2 = list2.Next
		} else {
			l.Val = list1.Val
			list1 = list1.Next
		}
	}
	return res.Next
}

// 递归调用
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists1(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists1(list2.Next, list1)
		return list2
	}
}
func TestMerge(t *testing.T) {

	tests := []struct {
		name    string
		list1   *ListNode
		list2   *ListNode
		wantRes []int
	}{
		{
			name:    "1",
			list1:   MakeList([]int{1, 2, 4}),
			list2:   MakeList([]int{1, 3, 4}),
			wantRes: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:    "2",
			list1:   MakeList([]int{1, 2, 4}),
			list2:   MakeList([]int{}),
			wantRes: []int{1, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MergeTwoLists(tt.list1, tt.list2)
			assert.Equal(t, tt.wantRes, res.Array())
		})
	}
}
