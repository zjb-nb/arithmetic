package array

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/4sum/description/
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	l := len(nums)
	for k, v := range nums {
		if v > target || k >= l-3 {
			return res
		}
		if k > 1 && nums[k] == nums[k-1] {
			continue
		}

		res = append(res, threeSumV1(nums[k+1:], target, v)...)

	}
	return res
}

func threeSumV1(nums []int, target int, val int) [][]int {
	res := [][]int{}
	l := len(nums)
	sort.Ints(nums)
	for cur, v := range nums {
		if v > target-val {
			return res
		}
		if cur > 1 && nums[cur] == nums[cur-1] {
			continue
		}
		left := cur + 1
		right := l - 1
		for left < right {
			if v+nums[left]+nums[right] == target-val {
				res = append(res, []int{val, v, nums[left], nums[right]})
				for left+1 < l && nums[left] == nums[left+1] {
					left++
				}
				left++
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if v+nums[left]+nums[right] > target-val {
				for right-1 > left && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else {
				for left+1 < l && nums[left] == nums[left+1] {
					left++
				}
				left++
			}
		}
	}
	return res
}

func Test_four_sum(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		target  int
		wantRes [][]int
	}{
		{
			name:   "[1,0,-1,0,-2,2],0",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			wantRes: [][]int{
				[]int{-2, -1, 1, 2},
				[]int{-2, 0, 0, 2},
				[]int{-1, 0, 0, 1},
			},
		},
		{
			name:   "[2,2,2,2],0",
			nums:   []int{2, 2, 2, 2},
			target: 8,
			wantRes: [][]int{
				[]int{2, 2, 2, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := fourSum(tt.nums, tt.target)
			assert.Equal(t, tt.wantRes, res)
		})
	}

}
