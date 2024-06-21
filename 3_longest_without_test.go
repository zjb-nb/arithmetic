package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://leetcode.cn/problems/longest-substring-without-repeating-characters/

/*
利用滑动窗口的思想
abcabcbb
以a开头:[abc]abccbb,[0,2]
以b开头:（0,2]区间范围检查过了，只需将右侧窗口继续扩大即可a[bca]bcbb
同理以c开头
*/
func lengthOfLongestSubstring(s string) int {
	//记录是否出现过该字节/字符
	m := make(map[byte]int)
	n := len(s)
	right, ans := -1, 0
	for i := 0; i < n; i++ {
		//前一个不再被需要，重置出现次数 a[bca]bcbb
		if i > 0 {
			delete(m, s[i-1])
		}
		//i为左侧窗口 right+1为移动右侧窗口
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]] = 1
			right++
		}
		ans = max(ans, right-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestLongest2(t *testing.T) {
	tests := []struct {
		name   string
		val    string
		wanRes int
	}{
		{
			name:   "1",
			val:    "abcabcbb",
			wanRes: 3,
		},
		{
			name:   "2",
			val:    "bbbb",
			wanRes: 1,
		},
		{
			name:   "3",
			val:    "pwwkew",
			wanRes: 3,
		},
		{
			name:   "4",
			val:    "",
			wanRes: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := lengthOfLongestSubstring(tt.val)
			assert.Equal(t, tt.wanRes, res)
		})
	}
}
