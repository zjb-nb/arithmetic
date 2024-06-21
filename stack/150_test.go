package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
// 逆波兰式
/*
review
*/
func evalrpn(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1, num2 := 0, 0
			if len(stack) > 0 {
				num1 = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				num2 = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			switch v {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}
	return stack[len(stack)-1]
}

func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		int_v, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, int_v)
		} else {
			prev, top := 0, 0
			if len(stack) > 0 {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 {
				prev = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			switch v {
			case "+":
				stack = append(stack, prev+top)
			case "-":
				stack = append(stack, prev-top)
			case "*":
				stack = append(stack, prev*top)
			case "/":
				stack = append(stack, prev/top)
			}
		}
	}
	return stack[len(stack)-1]
}

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}
	return stack[0]
}

func TestStack(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []string
		wantRes int
	}{
		{
			name:    "1",
			tokens:  []string{"2", "1", "+", "3", "*"},
			wantRes: 9,
		},
		{
			name:    "2",
			tokens:  []string{"4", "13", "5", "/", "+"},
			wantRes: 6,
		},
		{
			name:    "3",
			tokens:  []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			wantRes: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := evalrpn(tt.tokens)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
