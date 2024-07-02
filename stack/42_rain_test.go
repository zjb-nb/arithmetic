package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/trapping-rain-water/description/
计算柱状图的最大积雨面积
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
*/
func Trap(height []int) int {
	area := 0
	if len(height) <= 2 {
		return area
	}
	//找左边界和右边界
	/*
		1.单调栈，比栈顶元素低就入栈
		2.高就出栈，此时右边界为当前元素right=cur，出栈的栈顶元素为积水的底部top
		栈顶元素为左边界 left s = (min(right,left)-top)*(i-j-1)
		3.重复2 直到遇到高的或栈空
	*/
	stack := make([]int, len(height))
	for k, v := range height {
		for len(stack) > 0 && v > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				area += (Min(height[left], v) - top) * (k - left - 1)
			}
		}
		stack = append(stack, k)
	}
	return area
}

/*
左边界要去寻找left 高右边界right 即 h[right]>= h[left]  s-中间柱状图面积即可
*/
func trap(height []int) int {
	n := len(height)
	area := 0
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			right := i
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			area += (Min(height[right], height[left]) - height[top]) * (right - left - 1)
		}
		stack = append(stack, i)
	}
	return area
}

func TestRain(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes int
	}{
		{
			name:    "1",
			data:    []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			wantRes: 6,
		},
		{
			name:    "2",
			data:    []int{4, 2, 0, 3, 2, 5},
			wantRes: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Trap(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
