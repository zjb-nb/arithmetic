package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
给你一个整数数组 nums ，
请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），
返回其最大和。
子数组 是数组中的一个连续部分。
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func MaxSubArray(nums []int) int {
	//TODO 即求s[i]-s[j]=k的最大值 就是找最大的s[i]和最小的s[j]暴力求解
	// 找对于每个s[i]最小的s[j]即可
	/*
		s[0] = 0
		for xxxx {
		  s[xxx]=xxx
		}
		prev =0
		for xxx {
		  prev=xxx
		}
		for xxxx {
		  res = Max(s[i]=prev[i])
		}
	*/
	res := math.MinInt
	s := make([]int, len(nums)+1)
	prev := make([]int, len(s))
	for i, v := range nums {
		s[i+1] = s[i] + v
	}
	for i := 1; i < len(s); i++ {
		prev[i] = Min(prev[i-1], s[i])
	}
	for i := 1; i < len(s); i++ {
		res = Max(res, s[i]-prev[i-1])
	}
	return res
}

/*
1.暴力求解 s[i]-s[j]
2. 使得s[j]最小 ，那就再维护一个关于s的最小值
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := math.MinInt
	s := make([]int, len(nums)+1)
	for k, v := range nums {
		s[k+1] = s[k] + v
	}

	pre := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		pre[i] = Min(pre[i-1], s[i])
	}

	for i := 1; i <= len(nums); i++ {
		res = Max(res, s[i]-pre[i-1])
	}
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMaxSubArr(t *testing.T) {
	tests := []struct {
		name    string
		data    []int
		wantRes int
	}{
		{
			name:    "1",
			data:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			wantRes: 6,
		},
		{
			name:    "2",
			data:    []int{1},
			wantRes: 1,
		},
		{
			name:    "3",
			data:    []int{5, 4, -1, 7, 8},
			wantRes: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MaxSubArray(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
