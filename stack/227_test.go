package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculateii(s string) int {
	tokens, ops := []string{}, []string{}
	val := 0
	nums_start := false
	need_zero := true
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			val = val*10 + int(s[i]) - int('0')
			nums_start = true
			continue
		} else if nums_start {
			tokens = append(tokens, strconv.Itoa(val))
			val = 0
			nums_start = false
			need_zero = false
		}

		if s[i] == ' ' {
			continue
		}

		if s[i] == '(' {
			need_zero = true
			ops = append(ops, string(s[i]))
			continue
		}

		if s[i] == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != "(" {
				tokens = append(tokens, ops[len(ops)-1])
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
			need_zero = false
			continue
		}
		if need_zero {
			tokens = append(tokens, "0")
		}
		for len(ops) > 0 && getValueLevel(ops[len(ops)-1]) >= getValueLevel(string(s[i])) {
			tokens = append(tokens, ops[len(ops)-1])
			ops = ops[:len(ops)-1]
		}
		ops = append(ops, string(s[i]))
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
