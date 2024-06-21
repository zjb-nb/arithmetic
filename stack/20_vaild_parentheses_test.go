// https://leetcode.cn/problems/valid-parentheses/description/
package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
https://leetcode.cn/problems/valid-parentheses/
有效括号
输入：s = "()[]{}"
输出：true
*/
func IsValid(s string) bool {
	if len(s)%2 > 0 {
		return false
	}
	var symbolmap map[byte]byte = map[byte]byte{
		')': '(', ']': '[', '}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if symbolmap[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != symbolmap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}

/*
1.使用stack进行弹压栈
2.双指针
很明显栈的方式时间复杂度简单
go的栈很简单就是切片，push和pop就是划分切片
*/
func isValid(s string) bool {
	//快速失败
	if len(s)%2 == 1 {
		return false
	}
	var symbolmap map[byte]byte = map[byte]byte{
		')': '(', ']': '[', '}': '{',
	}
	//栈
	var stack []byte
	for i := 0; i < len(s); i++ {
		//命中的情况，尝试去栈顶弹出元素
		if symbolmap[s[i]] > 0 {
			//判断栈中有无元素且栈顶元素为对应括号
			//没有命中，快速失败
			if len(stack) == 0 || stack[len(stack)-1] != symbolmap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	//栈为空说明全部命中
	return len(stack) == 0
}

func TestXxx(t *testing.T) {
	tests := []struct {
		name    string
		val     string
		wantRes bool
	}{
		{
			name:    "()",
			val:     "()",
			wantRes: true,
		},
		{
			name:    "()[]",
			val:     "()[]",
			wantRes: true,
		},
		{
			name:    "([])",
			val:     "([])",
			wantRes: true,
		},
		{
			name:    "(",
			val:     "(",
			wantRes: false,
		},
		{
			name:    "([](",
			val:     "([](",
			wantRes: false,
		},
		{
			name:    "(}",
			val:     "(}",
			wantRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsValid(tt.val)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
