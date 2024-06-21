package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// review
// 跨台阶可以选择一次跨两步或者一步，求到达n级台阶的方法数量
/*
https://leetcode.cn/problems/climbing-stairs/description/
1.递归，2非递归
*/
func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}
	prev, cur := 1, 2
	for i := 3; i <= n; i++ {
		cur, prev = prev+cur, cur
	}
	return cur
}

/*
到达n阶台阶的方法数量取决于，到达n-1的数量+n-2的数量
*/
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	prev := 1
	now := 2
	for i := 3; i <= n; i++ {
		now, prev = now+prev, now
	}
	return now
}

func TestClimb(t *testing.T) {
	tests := []struct {
		name    string
		data    int
		wantRes int
	}{
		{
			name:    "a",
			data:    1,
			wantRes: 1,
		},
		{
			name:    "b",
			data:    3,
			wantRes: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ClimbStairs(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
