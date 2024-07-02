package main

import (
	"sort"
	"testing"
)

func PermuteUnique(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}
	sort.Ints(nums)
	/*
		递归每次插入一个，插入规则
		1.当len(tmp)==len(nums) 入栈结果
		2. for循环数组nums，我们要求插入的必须是之前没插入过的，且同一个插入位置不能出现相同的插入数字
		因此我们用visted表示该数字已经插入过了，if !visted[i] 即可插入
		3.解决同一个插入位置不能出现相同的插入数字
		j>0 && nums[j]==nums[j-1] ? 不行
		 && !visted[j-1] 即前面相同的已经用过了后面相同的不要了
		visted[j-1]=1 表示j-1正在插入，且j-1插入位置和j不相同
		反之=0 表示j-1插入过了，且当前插入位置和j-1相同
	*/
	n := len(nums)
	stack := make([]int, 0, n)
	visted := make([]int, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, append([]int{}, stack...))
			return
		}
		if i > n {
			return
		}
		for j := 0; j < n; j++ {
			if j > 0 && nums[j] == nums[j-1] && visted[j-1] == 0 {
				continue
			}
			if visted[j] == 0 {
				stack = append(stack, nums[j])
				visted[j] = 1
				dfs(i + 1)
				visted[j] = 0
				stack = stack[:len(stack)-1]
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
