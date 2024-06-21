package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/generate-parentheses/
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
*/
func GenerateParenthesis(n int) []string {
	res := []string{}
	var dfs func(left, right int, str string)
	dfs = func(left, right int, str string) {
		if left == 0 && right == 0 {
			res = append(res, str)
			return
		}
		if left > 0 {
			dfs(left-1, right, str+"(")
		}
		if right > 0 && right > left {
			dfs(left, right-1, str+")")
		}
	}
	dfs(n, n, "")
	return res
}

/*
添加左右括号的时机
*/
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	str := ""
	var add func(left, right int, str string)
	add = func(left, right int, str string) {
		if left == 0 && right == 0 {
			res = append(res, str)
			return
		}
		if right > left {
			add(left, right-1, str+")")
		}
		if left > 0 {
			add(left-1, right, str+"(")
		}
	}
	add(n, n, str)
	return res
}

// 这个测试用不了
func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		data    int
		wantRes []string
	}{
		// {
		// 	name:    "1",
		// 	data:    3,
		// 	wantRes: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		// },
		{
			name:    "2",
			data:    1,
			wantRes: []string{"()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := GenerateParenthesis(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
