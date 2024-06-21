package main

import "testing"

var m = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	res := []string{}
	if digits == "" {
		return res
	}
	str := []byte{}
	n := len(digits)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res = append(res, string(str))
			return
		}
		for j := 0; j < len(m[digits[i]]); j++ {
			str = append(str, m[digits[i]][j])
			dfs(i + 1)
			str = str[:len(str)-1]
		}
	}
	dfs(0)
	return res
}

func letterCombin(str string) (res []string) {
	n := len(str)
	if n == 0 {
		return res
	}
	var dfs func(i int)
	stack := []byte{}
	dfs = func(i int) {
		if i >= n {
			res = append(res, string(stack))
			return
		}

		temp := m[str[i]]
		for j := 0; j < len(temp); j++ {
			stack = append(stack, temp[j])
			dfs(i + 1)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)
	return
}

func TestLetter(t *testing.T) {
	t.Log(letterCombin("23"))
}
