package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review https://leetcode.cn/problems/basic-calculator/
基本计算器 根据字符串计算值,这里和227题加减乘除放一起
输入：s = "(1+(4+5+2)-3)+(6+8)"
输出：23
*/
func Calculate(s string) int {
	// 难点是处理 括号和 负数 负数可以看成(-a)
	/*
						2-1+2 [2,1,-,2+]  数字在入栈前检查 符号栈有无有就直接添
				    1+(1+1) [1,1,1,+,+]  遇到左括号入栈s2[+,(,+]，右括号出栈直到左括号 s1[1,1,1,+]  s2
						(1+(4+5+2)-3)+(6+8)  [01045+2++3-068++] [145+2++3-68+]
		         -1+(-1) [1-01-+]
	*/
	s1, s2 := make([]string, 0, len(s)), make([]string, 0, len(s))
	var need_zero bool = true
	for _, v := range s {
		if v == ' ' {
			continue
		}
		_, err := strconv.Atoi(string(v))
		if err == nil {
			s1 = append(s1, string(v))
			need_zero = false
			continue
		}
		val := string(v)
		if val == "(" {
			need_zero = true
			s2 = append(s2, string(v))
			continue
		}
		if val == ")" {
			for s2[len(s2)-1] != "(" {
				s1 = append(s1, s2[len(s2)-1])
				s2 = s2[:len(s2)-1]
			}
			s2 = s2[:len(s2)-1]
			continue
		}
		if need_zero {
			s1 = append(s1, "0")
			need_zero = false
		}
		if len(s2) > 0 && s2[len(s2)-1] != "(" && Level(val) <= Level(s2[len(s2)-1]) {
			s1 = append(s1, s2[len(s2)-1])
			s2 = s2[:len(s2)-1]
		}
		need_zero = true
		s2 = append(s2, string(v))
	}
	for len(s2) > 0 {
		s1 = append(s1, s2[len(s2)-1])
		s2 = s2[:len(s2)-1]
	}
	return eval(s1)
}

func Level(s string) int {
	switch s {
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	default:
		return 0
	}
}

// [2,1,+] = 3
func eval(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			prev, top := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, prev+top)
			case "-":
				stack = append(stack, prev-top)
			case "*":
				stack = append(stack, prev*top)
			case "/":
				stack = append(stack, prev/top)
			}
		} else {
			stack = append(stack, val)
		}
	}
	return stack[len(stack)-1]
}

/*
逆波兰式
1.遇到数字直接入tokens栈
2.遇到运算符，若前面ops有其他运算符，先把其他计算，再入ops栈本身
3.遇到右括号 进行ops出栈计算操作
4. 添加0 防止遇到(-1)这种情况
*/
func calculate(s string) int {
	ops := []string{}
	tokens := []string{}

	val := 0
	nums_start := false
	need_zero := true
	// byte是uint8的别名可以直接互相转换
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			val = val*10 + int(s[i]) - int('0')
			nums_start = true
			continue
		} else if nums_start {
			need_zero = false
			//此时说明前面已经找到了一个完整的数值，前面的数可以入栈了
			tokens = append(tokens, strconv.Itoa(val))
			val = 0
			nums_start = false
		}

		//处理当前非数值字符
		if s[i] == ' ' {
			continue
		}
		if s[i] == '(' {
			need_zero = true
			ops = append(ops, string(s[i]))
			continue
		}
		if s[i] == ')' {
			for ops[len(ops)-1] != "(" {
				tokens = append(tokens, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			// 将左括号弹出
			ops = ops[:len(ops)-1]
			need_zero = false
			continue
		}
		if need_zero {
			tokens = append(tokens, "0")
		}
		for len(ops) > 0 && ops[len(ops)-1] != "(" {
			tokens = append(tokens, ops[len(ops)-1])
			ops = ops[:len(ops)-1]
		}
		ops = append(ops, string(s[i]))
		//符号后面加0 是为了防止出现 -2+1的情况
		need_zero = true
	}
	if nums_start {
		tokens = append(tokens, strconv.Itoa(val))
	}
	for len(ops) > 0 {
		tokens = append(tokens, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return eval(tokens)
}

func TestC(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantRes int
	}{

		{
			name:    "1",
			data:    "1 + 1",
			wantRes: 2,
		},
		{
			name:    "2",
			data:    " 2-1 + 2 ",
			wantRes: 3,
		},
		{
			name:    "3",
			data:    "(1+(4+5+2)-3)+(6+8)",
			wantRes: 23,
		},
		{
			name:    "4",
			data:    "-1+2",
			wantRes: 1,
		},

		{
			name:    "5",
			data:    "(-1)+(-1)",
			wantRes: -2,
		},
		{
			name:    "6",
			data:    "3+2*2",
			wantRes: 7,
		},
		{
			name:    "7",
			data:    "3+2*2+1",
			wantRes: 8,
		},
		{
			name:    "8",
			data:    "3*(1+2*2+1)+(1+2*2+1)",
			wantRes: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Calculate(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
