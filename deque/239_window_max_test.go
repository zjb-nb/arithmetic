package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/sliding-window-maximum/description/
滑动窗口最大值
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
*/
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k == 1 {
		return nums
	}
	if k <= 0 {
		return []int{}
	}
	var add func(int)
	que := make([]int, 0, k)
	res := []int{}
	add = func(i int) {
		for len(que) > 0 && nums[i] > nums[que[len(que)-1]] {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}
	for i := 0; i < k; i++ {
		add(i)
	}
	res = append(res, nums[que[0]])
	for i := k; i < len(nums); i++ {
		if que[0] <= i-k {
			que = que[1:]
		}
		add(i)
		res = append(res, nums[que[0]])
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 1, n-k+1)
	que := []int{}
	push := func(i int) {
		for len(que) > 0 && nums[i] > nums[que[len(que)-1]] {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	res[0] = nums[que[0]]
	for i := k; i < n; i++ {
		push(i)
		for i-k >= que[0] {
			que = que[1:]
		}
		res = append(res, nums[que[0]])
	}

	return res
}

func TestMaxSlidingWinow(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		k       int
		wantRes []int
	}{
		{
			name:    "1",
			data:    []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:       3,
			wantRes: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:    "2",
			data:    []int{1, -1},
			k:       1,
			wantRes: []int{1, -1},
		},
		{
			name:    "3",
			data:    []int{1},
			k:       1,
			wantRes: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MaxSlidingWindow(tt.data, tt.k)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
