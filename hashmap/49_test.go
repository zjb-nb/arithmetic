package hashmap

import (
	"testing"
)

/*
https://leetcode.cn/problems/group-anagrams/
给你一个字符串数组，请你将 字母异位词 组合在一起。
可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/
func GroupAnagrams(strs []string) [][]string {
	/*
		1.用一个map存放键为索引，值为单词
	*/
	res := [][]string{}
	hmap := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := [26]int{}
		for i := 0; i < len(str); i++ {
			tmp[str[i]-'a']++
		}
		hmap[tmp] = append(hmap[tmp], str)
	}
	for _, v := range hmap {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	sums := map[[26]int][]string{}
	for _, str := range strs {
		c := [26]int{}
		for _, ch := range str {
			c[ch-'a']++
		}
		sums[c] = append(sums[c], str)
	}

	for _, v := range sums {
		res = append(res, v)
	}
	return res
}
func TestGroupanagrams(t *testing.T) {
	res := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	t.Log(res)
}
