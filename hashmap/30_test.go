package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
输出顺序无关紧要。返回 [9,0] 也是可以的。
*/
func FindSubstring(s string, words []string) []int {
	return []int{}
}

/*
"barfoo" mA=>["bar":1,foo:1]
["bar",foo],mB=>[bar:1,foo:1]
ma=?mb
*/
func findSubstring(s string, words []string) []int {
	m := len(words)
	n := len(words[0])
	if len(s) < m*n {
		return []int{}
	}
	hmap := map[string]int{}
	res := []int{}
	for _, v := range words {
		hmap[v]++
	}
	//"barfoothefoobarman"
	for i := 0; i+m*n <= len(s); i++ {
		my_hm := map[string]int{}
		//TODO 返回map
		for j := 0; j < m; j++ {
			my_hm[string(s[i+j*n:i+(j+1)*n])]++
		}
		//TODO map比较
		isVaild := true
		for str, v := range my_hm {
			if hmap[str] != v {
				isVaild = false
				break
			}
		}
		if isVaild {
			res = append(res, i)
		}
	}

	return res
}

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		words   []string
		wantRes []int
	}{
		{
			name:    "0",
			data:    "wordgoodgoodgoodbestword",
			words:   []string{"word", "good", "best", "good"},
			wantRes: []int{8},
		},
		{
			name:    "1",
			data:    "barfoothefoobarman",
			words:   []string{"foo", "bar"},
			wantRes: []int{0, 9},
		},
		{
			name:    "2",
			data:    "wordgoodgoodgoodbestword",
			words:   []string{"word", "good", "best", "word"},
			wantRes: []int{},
		},
		{
			name:    "3",
			data:    "barfoofoobarthefoobarman",
			words:   []string{"bar", "foo", "the"},
			wantRes: []int{6, 9, 12},
		},
		{
			name:    "3",
			data:    "lingmindraboofooowingdingbarrwingmonkeypoundcake",
			words:   []string{"fooo", "barr", "wing", "ding", "wing"},
			wantRes: []int{13},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := FindSubstring(tt.data, tt.words)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
