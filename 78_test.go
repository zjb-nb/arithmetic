package main

import "testing"

/*
https://leetcode.cn/problems/subsets/
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func Subsets(nums []int) [][]int {
	panic("")
}

/*
1.枚举子集的大小，再往里面填数字
k==n 停止，已经遍历过的不能再出现
*/
func subsets(nums []int) (res [][]int) {
	n := len(nums)
	stack := []int{}
	var dfs func(i, k int)

	dfs = func(i, k int) {
		if len(stack) == k {
			res = append(res, append([]int(nil), stack...))
			return
		}

		for j := i; j < n; j++ {
			stack = append(stack, nums[j])
			dfs(j+1, k)
			stack = stack[:len(stack)-1]
		}
	}
	for i := 1; i <= n; i++ {
		dfs(0, i)
	}
	return
}

/*
2.枚举数组中的每个元素加入或者不加入
*/
func subsets1(nums []int) (res [][]int) {
	n := len(nums)
	stack := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			if len(stack) == 0 {
				return
			}
			res = append(res, append([]int(nil), stack...))
			return
		}
		dfs(i + 1)
		stack = append(stack, nums[i])
		dfs(i + 1)
		stack = stack[:len(stack)-1]
	}
	dfs(0)
	return
}

/*
迭代
[[]]
[[],[1]]
[[],[1],[2],[1,2]]

*/
func subsets2(nums []int) [][]int {
	// n := len(nums)
	res := make([][]int, 1)
	for _, num := range nums {
		new_set := [][]int{}
		for _, set := range res {
			new_set = append(new_set, append([]int{num}, set...))
		}
		res = append(res, new_set...)
	}
	return res
}
func TestSubSet(t *testing.T) {
	t.Log(subsets2([]int{1, 2, 3}))
}
