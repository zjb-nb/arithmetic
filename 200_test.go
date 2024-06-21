package main

import "testing"

func numIslands(grid [][]byte) int {
	count := 0
	len_row := len(grid)
	len_col := len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row > len_row-1 || col < 0 || col > len_col-1 || grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col-1)
		dfs(row, col+1)

	}
	for k, row := range grid {
		for col, v := range row {
			if v == '1' {
				count++
				dfs(k, col)
			}
		}
	}

	return count
}

func TestSumIsLand(t *testing.T) {
	t.Log(numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}))
}
