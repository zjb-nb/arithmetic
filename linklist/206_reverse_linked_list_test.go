package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
recover
递归 https://leetcode.cn/problems/reverse-linked-list/solutions/36710/dong-hua-yan-shi-206-fan-zhuan-lian-biao-by-user74/
[1,2,3,4,5]
[5,4,3,2,1]
2种写法
*/
func ReverseList(head *ListNode) *ListNode {
	panic("")
}

/*
https://leetcode.cn/problems/reverse-linked-list/
反单转链表
*/

/*
1->2 -> nil
nil<-1<-2
prev为前一个，因为是单向链表没存储前一个，需要手动存储
cur为当前
next为下一个
*/
func reverseList1(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		//推动循环前进，存储前一个
		prev = curr
		curr = next
	}
	return prev
}

/*
方法3 利用栈的先入后出

*/

/*
方法4 递归
*/
func reverseList4(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList4(head.Next)
	head.Next.Next = head
	//单向链表，防止成环
	head.Next = nil
	return newHead
}

func TestRever(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes []int
	}{
		{
			name:    "[1,2]",
			data:    []int{1, 2},
			wantRes: []int{2, 1},
		},
		{
			name:    "[1,2,3,4,5]",
			data:    []int{1, 2, 3, 4, 5},
			wantRes: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := MakeList(tt.data)
			res := ReverseList(list).Array()
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
