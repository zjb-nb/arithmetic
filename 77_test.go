package main

import (
	"testing"
)

/*
https://leetcode.cn/problems/combinations/
输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
*/
func Combine(n int, k int) (res [][]int) {
	/*
		迭代，每次投入一个
	*/
	if n < k || k <= 0 {
		return
	}
	tmp := make([]int, 0, k)
	var dfs func(m int)
	dfs = func(m int) {
		if len(tmp) == k {
			res = append(res, append([]int{}, tmp...))
			return
		}
		if m > n {
			return
		}
		for i := m; i <= n; i++ {
			tmp = append(tmp, i)
			dfs(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(1)
	return
}

func combine(n int, k int) (res [][]int) {
	stack := []int{}
	var add func(i int)
	add = func(i int) {
		if len(stack) == k {
			res = append(res, append([]int(nil), stack...))
			return
		}
		if i > n {
			return
		}
		for j := i; j <= n; j++ {
			stack = append(stack, j)
			add(j + 1)
			stack = stack[:len(stack)-1]
		}
	}
	add(1)
	return
}

/*
arr := []int{1, 3}
res := [][]int{}
arr = append(arr, 4)
res = append(res, arr)
t.Log(res)
arr = arr[:len(arr)-1]
arr = append(arr, 5)
t.Log(res)
*/
func combineBug(n int, k int) (res [][]int) {
	if n < k {
		return
	}
	var dfs func(temp []int, i int)
	dfs = func(temp []int, i int) {
		if len(temp) == k {
			res = append(res, temp)
			return
		}
		if i > n {
			return
		}
		for j := i; j <= n; j++ {
			arr := temp
			arr = append(arr, j)
			dfs(arr, j+1)
		}
	}
	dfs([]int{}, 1)
	return res
}

func TestMain(t *testing.T) {
	t.Log(Combine(4, 2))
}
