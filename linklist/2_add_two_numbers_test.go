package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewList(input []int) *ListNode {
	if input == nil {
		return nil
	}
	f := &ListNode{Val: input[0]}
	if len(input) == 1 {
		return f
	}
	pre := f
	for _, v := range input[1:] {
		pre.Next = &ListNode{
			Val: v,
		}
		pre = pre.Next
	}
	return f
}

func Print(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

/*
465+342=807
465: 5->6->4
342: 2->4->3
807: 7->0->8
进位情况？多加一个
92+9=101
92：2->9
9   9
101:1->0->1
很明显停止条件就是 l1==nil && l2==nil
时间On
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	l := res
	tmp := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}

		l1_val, l2_val := 0, 0
		if l1 != nil {
			l1_val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2_val = l2.Val
			l2 = l2.Next
		}

		if l == nil {
			res = &ListNode{}
			l = res
		} else {
			l.Next = &ListNode{}
			l = l.Next
		}
		l.Val = (l1_val + l2_val + tmp) % 10
		tmp = (l1_val + l2_val + tmp) / 10

	}

	if tmp > 0 {
		l.Next = &ListNode{Val: 1}
	}
	return res
}

func TestListNode(t *testing.T) {
	l := NewList([]int{2, 4, 3})
	t.Log(Print(l))
}

func TestAddTwoNumbers(t *testing.T) {
	test := []struct {
		name    string
		l1      []int
		l2      []int
		wantRes []int
	}{
		{
			name:    "first",
			l1:      []int{2, 4, 3},
			l2:      []int{5, 6, 4},
			wantRes: []int{7, 0, 8},
		},
		{
			name:    "second",
			l1:      []int{0},
			l2:      []int{0},
			wantRes: []int{0},
		}, {
			name:    "third",
			l1:      []int{9, 9, 9, 9, 9, 9, 9},
			l2:      []int{9, 9, 9, 9},
			wantRes: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewList(tt.l1)
			l2 := NewList(tt.l2)
			res := addTwoNumbers(l1, l2)
			int_res := Print(res)
			assert.Equal(t, tt.wantRes, int_res)
		})
	}
}
