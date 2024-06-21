package main

import "testing"

var dirX = []int{0, 1, 1, 1, 0, -1, -1, -1}
var dirY = []int{1, 0, 1, -1, -1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	var dfs func(row, col int)
	len_row, len_col := len(board), len(board[0])
	dfs = func(x, y int) {
		cnt := 0
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx > len_row-1 || ty < 0 || ty > len_col-1 {
				continue
			}
			if board[tx][ty] == 'M' {
				cnt++
			}
		}
		if cnt > 0 {
			board[x][y] = byte(cnt + '0')
		} else {
			board[x][y] = 'B'
			for i := 0; i < 8; i++ {
				tx, ty := x+dirX[i], y+dirY[i]
				if tx < 0 || tx > len_row-1 || ty < 0 || ty > len_col-1 || board[tx][ty] != 'E' {
					continue
				}
				dfs(tx, ty)
			}
		}
	}
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
	} else {
		dfs(click[0], click[1])
	}
	return board
}

func TestBoard(t *testing.T) {
	t.Log(updateBoard([][]byte{
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'M', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
		[]byte{'E', 'E', 'E', 'E', 'E'},
	}, []int{3, 0}))
}
