package main

import "testing"

//https://leetcode.cn/problems/check-distances-between-same-letters/submissions/422791655/

/*
只需要记住他第一次出现的下标即可，利用map存储
第二次出现就和distance检查对比即可
*/
func checkDistances(s string, distance []int) bool {
	var str_map map[byte]int = make(map[byte]int)
	//在map中就跳过，分别计算数量
	for i := 0; i < len(s); i++ {
		if v, ok := str_map[s[i]]; ok {
			// str_map[s[i]] = i-v-1
			//存在就对比
			if distance[s[i]-97] != i-v-1 {
				return false
			}
		} else {
			str_map[s[i]] = i
		}
	}
	return true
}

func TestCheckDistances(t *testing.T) {
	t.Log(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	t.Log(checkDistances("aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
