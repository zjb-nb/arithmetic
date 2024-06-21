package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://leetcode.cn/problems/roman-to-integer/

/*
后一个与前一个比较，比他大则进行减法操作
时间复杂度On
空间复杂度O1
*/
func romanToInt(s string) int {
	var res, prev, tmp int
	str_byte := []byte(s)
	for k, v := range str_byte {
		tmp = change(v)
		if k-1 < 0 {
			res += tmp
			prev = tmp
			continue
		}

		//比较
		if tmp > prev {
			res = res - 2*prev + tmp // res-prev+tmp-prev
		} else {
			res += tmp
		}

		prev = tmp
	}
	return res
}

// 精简一下第一步的代码
func romanToInt2(s string) int {
	var res, prev, tmp int
	str_byte := []byte(s)
	for k, v := range str_byte {
		tmp = change(v)
		if k-1 >= 0 {
			if tmp > prev {
				res = res - 2*prev
			}
		}

		res += tmp
		prev = tmp
	}
	return res
}

// 使用map，并且与右侧数字进行对比，不是左侧
func romanToInt3(s string) int {
	var res, tmp int
	str_byte := []byte(s)
	for k, v := range str_byte {
		tmp = symbolValues[v]
		if k < len(str_byte)-1 && tmp < symbolValues[str_byte[k+1]] {
			res -= tmp
		} else {
			res += tmp
		}
	}
	return res
}

func change(s byte) int {
	var res int
	switch s {
	case 'I':
		res = 1
	case 'V':
		res = 5
	case 'X':
		res = 10
	case 'L':
		res = 50
	case 'C':
		res = 100
	case 'D':
		res = 500
	case 'M':
		res = 1000
	}
	return res
}

// 用map替代change
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func TestRoman(t *testing.T) {
	tests := []struct {
		name    string
		val     string
		wantRes int
	}{
		{
			name:    "III",
			val:     "III",
			wantRes: 3,
		},
		{
			name:    "IV",
			val:     "IV",
			wantRes: 4,
		},
		{
			name:    "IX",
			val:     "IX",
			wantRes: 9,
		},
		{
			name:    "LVIII",
			val:     "LVIII",
			wantRes: 58,
		},
		{
			name:    "MCMXCIV",
			val:     "MCMXCIV",
			wantRes: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := romanToInt3(tt.val)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
