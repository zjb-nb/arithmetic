package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//https://leetcode.cn/problems/palindrome-number/
/*
X   正数，回文即对称,左右两个标游走对比，相遇即true
时间o logn
空间复杂度很大
*/
func isPalindrome(x int) bool {
	//快速返回
	// 小于0一定不是回文
	// 同理个位为0也不是回文，开头不能为0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 个位数字一定为回文
	if x < 10 {
		return true
	}
	input := []rune(strconv.Itoa(x))
	right := len(input) - 1
	left := 0
	for left <= right {
		//为个位数
		if left == right {
			return true
		}
		if input[left] == input[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

/*
完全反转
时间复杂度为O1
*/
func isPalindrome2(x int) bool {
	//快速返回
	// 小于0一定不是回文
	// 同理个位为0也不是回文，开头不能为0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 个位数字一定为回文
	if x < 10 {
		return true
	}

	res := 0
	target := x
	for target > 0 {
		res = res*10 + target%10
		target /= 10
	}

	return res == x
}

/*
翻转一半 即 1234 ->  前一个12和 后一个43做对比，即看这张纸能否对折
没翻转完的时候 前一个一定比后一个大
偶数情况
1234
123 > 4
12  < 43 停
1221
122 >1
12  = 12 停
寄数情况
123
12 > 3
1  < 32 停，最后一个为中位数舍去 x/10
时间复杂度o logn
空间复杂度 o1 只有一个常量
*/
func isPalindrome3(x int) bool {
	//快速返回
	// 小于0一定不是回文
	// 同理个位为0也不是回文，开头不能为0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 个位数字一定为回文
	if x < 10 {
		return true
	}

	target := 0
	for x > target {
		target = target*10 + x%10
		x /= 10
	}

	return x == target || x == (target/10)
}

func TestPalindorme(t *testing.T) {
	tests := []struct {
		name    string
		val     int
		wantRes bool
	}{
		{
			name:    "first",
			val:     121,
			wantRes: true,
		},
		{
			name:    "second",
			val:     -121,
			wantRes: false,
		},
		{
			name:    "third",
			val:     10,
			wantRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := isPalindrome(tt.val)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestP3(t *testing.T) {
	t.Log(isPalindrome3(121))
	t.Log(isPalindrome3(10))
	t.Log(isPalindrome3(123))
	t.Log(isPalindrome3(1221))
}
