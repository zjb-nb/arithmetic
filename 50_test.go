package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/powx-n/
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
*/
func Pow(x float64, n int) float64 {
	/*
		迭代 n>0 x*x*x n<0 1/x * 1/x *1/x
		将n一分为2进行计算 奇数就是再乘以本身 当n<0 n%2==-1 所以只以偶数情况分类讨论
	*/
	if n == 0 {
		return float64(1)
	}
	y := x
	if n < 0 {
		y = 1 / x
	}
	if n%2 == 0 {
		return Pow(x, n/2) * Pow(x, n/2)
	}
	return Pow(x, n/2) * Pow(x, n/2) * y
}

func pow(x float32, n int) float32 {
	if n == 0 {
		return float32(1)
	}
	y := x
	if n < 0 {
		y = 1 / x
	}

	if n%2 == 0 {
		return pow(x, n/2) * pow(x, n/2)
	}

	return pow(x, n/2) * pow(x, n/2) * y
}

func TestPow(t *testing.T) {
	tests := []struct {
		name    string
		x       float64
		n       int
		wantRes float64
	}{
		{
			name:    "1",
			x:       2.00000,
			n:       10,
			wantRes: 1024.00000,
		},
		{
			name:    "2",
			x:       2.10000,
			n:       3,
			wantRes: 9.26100,
		},
		{
			name:    "1",
			x:       2.00000,
			n:       -2,
			wantRes: 0.25000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Pow(tt.x, tt.n)
			assert.Equal(t, int64(tt.wantRes*1000), int64(res*1000))
		})
	}
}
