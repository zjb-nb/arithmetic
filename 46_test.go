package main

import "testing"

func permute(nums []int) (res [][]int) {

	n := len(nums)
	stack := make([]int, n)
	var dfs func(i int)
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int(nil), stack...))
			return
		}
		for j := 0; j < n; j++ {
			if m[nums[j]] > 0 {
				stack[i] = nums[j]
				m[nums[j]]--
				dfs(i + 1)
				m[nums[j]]++
			}
		}
	}
	dfs(0)
	return
}

func TestPermute(t *testing.T) {
	t.Log(permute([]int{1, 2, 3}))
}
