package main

import (
	"testing"
)

func solveNQueens(n int) [][]string {
	pie := make(map[int]int)
	na := make(map[int]int)
	col := make([]int, n)
	var dfs func(row int)
	stack := []int{}
	res := [][]string{}
	dfs = func(row int) {
		if row == n {
			res = append(res, makeRes(stack))
			return
		}

		for i := 0; i < n; i++ {
			if col[i] > 0 || na[i-row] > 0 || pie[i+row] > 0 {
				continue
			}
			col[i] = 1
			na[i-row] = 1
			pie[i+row] = 1
			stack = append(stack, i)
			dfs(row + 1)
			col[i] = 0
			na[i-row] = 0
			pie[i+row] = 0
			stack = stack[:len(stack)-1]
		}
	}

	dfs(0)
	return res
}

/*
n皇后，8个方向不能攻击
1.放入前判断是否会被攻击
2. 空间换时间，放入的同时，放入攻击范围 列、撇、捺
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
