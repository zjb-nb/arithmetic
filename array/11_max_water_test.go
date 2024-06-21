package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/container-with-most-water/description/
// review [1,8,6,2,5,4,8,3,7]
// s= h * w = min(height[i],height[j])*(height[j]-height[i])
func MaxArea(height []int) int {
	s, l, r := 0, 0, len(height)-1
	for l < r {
		if height[r] < height[l] {
			s = Max(s, height[r]*(r-l))
			r--
		} else {
			s = Max(s, height[l]*(r-l))
			l++
		}
	}
	return s
}

// [1,8,6,2,5,4,8,3,7]
// 方法1.强制计算出每一个面积
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			c := (j - i) * Min(height[j], height[i])
			max = Max(max, c)
		}
	}
	return max
}

/*
方法2 向内移动短板，计算面积--双指针
area := min(h[i],h[j])*( j-i )
1. 向内移动短板，min(h[i],h[j])可能增大，面积可能增大
2.        长    （j-i）一定减少，面积一定减少
*/
func maxArea2(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max, i, j := 0, 0, len(height)-1
	for i < j {
		if height[j] > height[i] {
			max = Max(max, (j-i)*height[i])
			i++
		} else {
			max = Max(max, (j-i)*height[j])
			j--
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes int
	}{
		{
			name:    "a",
			data:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			wantRes: 49,
		},
		{
			name:    "b",
			data:    []int{},
			wantRes: 0,
		},
		{
			name:    "c",
			data:    []int{2},
			wantRes: 0,
		},
		{
			name:    "d",
			data:    []int{1, 1},
			wantRes: 1,
		},
	}
	for _, tt := range tests {
		t.Run("a", func(t *testing.T) {
			res := MaxArea(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
