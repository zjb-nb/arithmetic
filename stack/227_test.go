package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculateii(s string) int {
	/*
			1. 目标就是转为逆波兰式
			2. 遇到数字就把数字添加tokens，遇到op就添加到ops
			3. op添加到ops时进行优先级判断，当前比栈顶低就把栈顶的弹出去，左括号‘（’直接入栈，
			右括号‘）’ 出栈直到遇到左括号
			4. 处理正负数，-1，引入need_zero ，一开始为tue，-1+2 =》 0-1+2
			计算符号op 改为true
			( 时为true (-1)
			 数字入栈后 为false
		  )为false
			5. 最后ops全部出栈
			伪代码如下
			zero = true
			for k,v :=range s {
			  if ==token {
				   tokens.append
					 zero = false
				}

				if == ' ' {continue}
				if == ( { ops.push ,zero=true}
				if == ) {ops.pop,zero=false}

				if zero {tokens.append(0)}

				if len(ops)>0 && now_level < stack_top_level {
				   ops.pop
				}

				ops.append

			}
				for ops {
				  tokens.append
				}
	*/
	tokens, ops := make([]string, 0, len(s)), make([]string, 0, len(s))
	need_zero, num_start := true, false
	token := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			num_start = true
			val, _ := strconv.Atoi(string(s[i]))
			token = token*10 + val
			continue
		} else if num_start {
			tokens = append(tokens, strconv.Itoa(token))
			token = 0
			need_zero = false
			num_start = false
		}
		str := string(s[i])
		if str == "(" {
			ops = append(ops, str)
			need_zero = true
			continue
		}
		if str == ")" {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				tokens = append(tokens, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
			// need_zero = false
			continue
		}

		if need_zero {
			tokens = append(tokens, "0")
		}

		if len(ops) > 0 && getValueLevel(ops[len(ops)-1]) >= getValueLevel(str) {
			tokens = append(tokens, ops[len(ops)-1])
			ops = ops[:len(ops)-1]
		}
		ops = append(ops, str)
		need_zero = true
	}
	if need_zero {
		tokens = append(tokens, strconv.Itoa(token))
	}
	for len(ops) > 0 {
		tokens = append(tokens, ops[len(ops)-1])
		ops = ops[:len(ops)-1]
	}
	return eval(tokens)
}

func getValueLevel(s string) int {
	if s == "+" || s == "-" {
		return 1
	}
	if s == "*" || s == "/" {
		return 2
	}
	return 0
}

func TestCii(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantRes int
	}{
		{
			name:    "1",
			data:    "3+2*2",
			wantRes: 7,
		},
		{
			name:    "2",
			data:    " 3/2 ",
			wantRes: 1,
		},
		{
			name:    "3",
			data:    " 3+5 / 2 ",
			wantRes: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := calculateii(tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
