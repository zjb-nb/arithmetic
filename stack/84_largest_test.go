package stack

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
求柱状图形成的最大矩形面积
输入：heights = [2,1,5,6,2,3]
输出：10
*/
func LargestRectangleArea(heights []int) int {
	s := 0
	//存下表
	heights = append(heights, -1)
	stack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			right := i - 1
			left := -1
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && h == heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			s = Max(s, h*(right-left))
		}
		stack = append(stack, i)
	}
	return s
}

/*
单调栈解法
h * (right-left)
1.当栈为空时，或入栈元素大于栈顶元素入栈
2.栈底表示柱状图的左边界即-1，栈保存下标
3.当前元素i小于栈顶元素cur，表示以栈顶柱子高度计算面积的矩形，能达到的最远右边界为 i-1
4. 弹出cur，当cur>栈顶元素top，表示能达到的最远左边界为top
此时能计算cur为高的矩形面积
*/
func largestRectangleArea(heights []int) int {
	area := 0
	heights = append(heights, 0)
	n := len(heights)
	stack := []int{}
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			temp := heights[stack[len(stack)-1]]
			right := i - 1
			left := -1
			stack = stack[:len(stack)-1]
			for len(stack) > 0 && temp == heights[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			area = Max(area, temp*(right-left))
		}
		stack = append(stack, i)
	}
	return area
}

// 枚举宽
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	area := 0
	for i := 0; i < n; i++ {
		minh := math.MaxInt
		for j := i; j < n; j++ {
			minh = Min(minh, heights[j])
			area = Max(area, minh*(j-i+1))
		}
	}
	return area
}

// 枚举高
func largestRectangleArea2(heights []int) int {
	n := len(heights)
	area := 0
	for i := 0; i < n; i++ {
		h := heights[i]
		left, right := i, i
		for left-1 >= 0 && heights[left-1] >= h {
			left--
		}
		for right+1 < n && heights[right+1] >= h {
			right++
		}
		area = Max(area, h*(right-left+1))
	}
	return area
}

func TestM(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes int
	}{
		{
			name:    "1",
			data:    []int{2, 1, 5, 6, 2, 3},
			wantRes: 10,
		},
		{
			name:    "2",
			data:    []int{0, 9},
			wantRes: 9,
		},
		{
			name:    "2",
			data:    []int{2, 4},
			wantRes: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LargestRectangleArea(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
