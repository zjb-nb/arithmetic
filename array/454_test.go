package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/4sum-ii/description/
// 4元组之和为0出现的次数
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	h := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			h[v1+v2]++
		}
	}
	for _, v1 := range nums3 {
		for _, v2 := range nums4 {
			count += h[-v1-v2]
		}
	}
	return count
}
func Test_fourV1_sum(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int

		wantRes int
	}{
		{
			name:  "1",
			nums1: []int{1, 2},
			nums2: []int{-2, -1},
			nums3: []int{-1, 2},
			nums4: []int{0, 2},

			wantRes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fourSumCount(tt.nums1, tt.nums2, tt.nums3, tt.nums4)
			assert.Equal(t, tt.wantRes, res)
		})
	}

}
