package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/range-sum-query-2d-immutable/
二维区域和检索 - 矩阵不可变
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1)
，右下角 为 (row2, col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

*/

/*
i 为行，
*/
type NumMatrix struct {
	sums [][]int
}

// s[i,j] = s[i-1,j]+s[j-1,i]-s[i-1.j-1]+nums[i,j]
func MakeNumMatrix(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	sums := make([][]int, m+1)
	sums[0] = make([]int, n+1)
	for i, row := range matrix {
		sums[i+1] = make([]int, n+1)
		for j, v := range row {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + v
		}
	}
	return NumMatrix{sums: sums}
}

// row2>row1 ,row2表示右下角
// s[row1,col1] -> s[row2,col2] = s[row2+1,col2+1] - s[row1,col2+1]-s[row2+1,col1]+s[row1,col1]
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] -
		this.sums[row1][col2+1] -
		this.sums[row2+1][col1] + this.sums[row1][col1]
}

func TestMatrix(t *testing.T) {
	m := MakeNumMatrix([][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	})
	assert.Equal(t, 8, m.SumRegion(2, 1, 4, 3))
	assert.Equal(t, 11, m.SumRegion(1, 1, 2, 2))
	assert.Equal(t, 12, m.SumRegion(1, 2, 2, 4))

}
