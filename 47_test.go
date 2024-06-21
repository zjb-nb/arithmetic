package main

import (
	"sort"
	"testing"
)

func PermuteUnique(nums []int) (res [][]int) {
	sort.Ints(nums)
	stack := []int{}
	n := len(nums)
	visted := make([]bool, n)
	var dfs func(i int)
	dfs = func(i int) {
		if len(stack) == n {
			res = append(res, append([]int(nil), stack...))
			return
		}

		for j := 0; j < n; j++ {
			if j > 0 && nums[j] == nums[j-1] && !visted[j-1] {
				continue
			}
			if !visted[j] {
				visted[j] = true
				stack = append(stack, nums[j])
				dfs(i + 1)
				stack = stack[:len(stack)-1]
				visted[j] = false
			}
		}
	}
	dfs(0)
	return
}
func permuteUnique(nums []int) (res [][]int) {
	n := len(nums)
	stack := []int{}
	sort.Ints(nums)
	not_visted := make([]bool, n)

	var dfs func(i int)
	dfs = func(i int) {
		if len(stack) == n {
			res = append(res, append([]int(nil), stack...))
			return
		}
		if i >= n {
			return
		}
		for j := 0; j < n; j++ {
			if j > 0 && nums[j] == nums[j-1] && not_visted[j-1] == false {
				continue
			}
			if !not_visted[j] {
				stack = append(stack, nums[j])
				not_visted[j] = true
				dfs(i + 1)
				not_visted[j] = false
				stack = stack[:len(stack)-1]
			}
		}
	}
	dfs(0)
	return res
}
func TestPermuteUnique(t *testing.T) {
	t.Log(PermuteUnique([]int{1, 1, 2}))
}
