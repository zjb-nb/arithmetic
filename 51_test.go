package main

import (
	"testing"
)

func solveNQueens(n int) [][]string {
	/*
		1.放之前检查有无冲突 即横竖撇捺 横竖很好判断，问题是撇捺
		2.
		stack存放每一行 存放的列数 [1,2,3] 如stack[0] 表示放在第0行的第1列
		dfs(j)
		for i range n  {
		   for k,v range stack {
			     if k==i {break}
			     if k-i == v-j || k-i == j-v {break}
			 }

		}

	*/
	res := make([][]string, 0, 4)
	stack := make([]int, 0, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, makeRes(stack))
			return
		}

		//枚举列数
		for j := 0; j < n; j++ {
			vaild := true
			for k, v := range stack {
				if v == j {
					vaild = false
					break
				}
				if j-v == k-i || j-v == i-k {
					vaild = false
					break
				}
			}
			if vaild {
				stack = append(stack, j)
				dfs(i + 1)
				stack = stack[:len(stack)-1]
			}
		}
	}
	dfs(0)
	return res

}

/*
n皇后，8个方向不能攻击
1.放入前判断是否会被攻击
2. 空间换时间，放入的同时，放入攻击范围 列、撇、捺
[0,0][0,1][0,2][0,3]
[1,0][1,1][1,2][1,3]
[2,0][2,1][2,2][2,3]
[3,0][3,1][3,2][3,3]
*/
func sloveQueen(n int) [][]string {
	res := [][]string{}
	stack := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, makeRes(stack))
			return
		}
		//随便放，然后看有没有冲突
		for j := 0; j < n; j++ {
			vaild := true
			for k, v := range stack {
				if v == j {
					vaild = false
					break
				}
				if i-k == j-v || i-k == v-j {
					vaild = false
					break
				}
			}
			if vaild {
				stack = append(stack, j)
				dfs(i + 1)
				stack = stack[:len(stack)-1]
			}
		}
	}
	dfs(0)
	return res
}

func makeRes(nums []int) []string {
	res := []string{}
	n := len(nums)
	for _, v := range nums {
		str := ""
		for i := 0; i < n; i++ {
			if v == i {
				str += "Q"
			} else {
				str += "."
			}
		}
		res = append(res, str)
	}
	return res
}

func TestSloveQueen(t *testing.T) {
	t.Log(solveNQueens(4))
}
