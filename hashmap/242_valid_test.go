package hashmap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/valid-anagram/description/
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

输入: s = "anagram", t = "nagaram"
输出: true
*/
func IsAnagram(s string, t string) bool {
	return true
}

/*
1.排序看是否相同
2.哈希表，一个加一个减
3.哈希表的进阶，字符串只包含26个字母，因此可以将哈希表变为数组
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hmap := map[rune]int{}
	for _, v := range s {
		hmap[v]++
	}
	for _, v := range t {
		hmap[v]--
	}
	for _, v := range hmap {
		if v > 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	return string(s1) == string(s2)
}

func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

func TestVaild(t *testing.T) {
	tests := []struct {
		name    string
		s1      string
		s2      string
		wantRes bool
	}{
		{
			name:    "1",
			s1:      "anagram",
			s2:      "nagaram",
			wantRes: true,
		},
		{
			name:    "1",
			s1:      "rat",
			s2:      "car",
			wantRes: false,
		},
		{
			name:    "1",
			s1:      "",
			s2:      "nagaram",
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsAnagram(tt.s1, tt.s2)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestSort(t *testing.T) {
	s1 := []byte("anagram")
	s2 := []byte("nagaram")
	sort.SliceStable(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.SliceStable(s2, func(i, j int) bool { return s2[i] < s2[j] })
	t.Log(string(s1), string(s2))
}
