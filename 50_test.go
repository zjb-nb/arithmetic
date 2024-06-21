package main

import "testing"

/*
https://leetcode.cn/problems/powx-n/
实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return Pow(x, n)
}

func Pow(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 1 {
		return Pow(x, n/2) * Pow(x, n/2) * x
	}
	return Pow(x, n/2) * Pow(x, n/2)
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
	t.Log(pow(2.000, 10))
	t.Log(pow(2.100, 3))
	t.Log(pow(2.000, -2))

}
