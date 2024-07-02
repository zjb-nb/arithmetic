package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/longest-common-prefix/
func LongestCommonPrefix(strs []string) string {
	/*
		将大问题划分为小问题 分治 左右区间的最小前缀

	*/
	var dfs func(left, right int) string
	dfs = func(left, right int) string {
		if left == right {
			return strs[left]
		}
		if left > right {
			return ""
		}
		mid := (left + right) / 2
		l_str := dfs(left, mid)
		r_str := dfs(mid+1, right)
		min_len := Min(len(l_str), len(r_str))
		i := 0
		for i < min_len {
			if l_str[i] != r_str[i] {
				break
			}
			i++
		}
		return l_str[:i]
	}
	return dfs(0, len(strs)-1)
}

/*
纵向扫描
不断对比，相同则移动指针
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	var longest int
	var tmp byte
	target := strs[0]
	for i := 0; i < len(target); i++ {
		tmp = target[i]
		for _, v := range strs {
			if i > len(v)-1 {
				return target[:longest]
			}
			if tmp != v[i] {
				return target[:longest]
			}
		}
		longest += 1
	}

	return target[:longest]
}

/*
分治，不断划分连个区间去寻找最短前缀，分解为子问题
*/
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func TestLongest(t *testing.T) {
	tests := []struct {
		name    string
		datas   []string
		wantRes string
	}{
		{
			name:    "1",
			datas:   []string{"flower", "flow", "flight"},
			wantRes: "fl",
		},
		{
			name:    "2",
			datas:   []string{"dog", "racecar", "car"},
			wantRes: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LongestCommonPrefix(tt.datas)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
