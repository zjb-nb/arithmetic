package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/plus-one/
*/
func PlusOne(digits []int) []int {
	res := []int{1}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res = append(res, digits...)
	return res
}

/*
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
*/
func plusOne(digits []int) []int {
	for n := len(digits) - 1; n >= 0; n-- {
		if digits[n] != 9 {
			digits[n]++
			return digits
		}
		digits[n] = 0
	}
	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}

func TestMy(t *testing.T) {
	tests := []struct {
		data    []int
		wantRes []int
	}{
		{
			data:    []int{1, 2, 3},
			wantRes: []int{1, 2, 4},
		},
		{
			data:    []int{1, 9, 9},
			wantRes: []int{2, 0, 0},
		},
		{
			data:    []int{9, 9},
			wantRes: []int{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run("go ", func(t *testing.T) {
			res := PlusOne(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
