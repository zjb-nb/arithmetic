package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 用递归 1221
func isPalindrome(head *ListNode) bool {
	cur := head
	var reverse func(c *ListNode) bool
	reverse = func(c *ListNode) bool {
		if c != nil {
			if !reverse(c.Next) {
				return false
			}
			if c.Val != cur.Val {
				return false
			}
			cur = cur.Next
		}
		return true
	}
	return reverse(cur)
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		n       *ListNode
		wantRes bool
	}{
		{
			name:    "1",
			n:       &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			wantRes: true,
		},
		{
			name:    "2",
			n:       &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			wantRes: false,
		},
		{
			name:    "3",
			n:       (*ListNode)(nil),
			wantRes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.n)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
