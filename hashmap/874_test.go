package hashmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
review
模拟机器人走路
command:
-1 右转
-2 左转
+X 向前移动X
obstacles为障碍物地点，如果遇到障碍物会停止前面，然后执行下一步指令
初始方向为 +Y轴
返回所到地点的最大欧氏距离 即最远点到原点的距离 X^2
*/
func RobotSim(commands []int, obstacles [][]int) int {
	res := 0
	directions := [][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}
	hmap := map[[2]int]int{}
	for _, v := range obstacles {
		hmap[[2]int{v[0], v[1]}]++
	}
	dire, curx, cury := 0, 0, 0
	for _, v := range commands {
		if v == -1 {
			dire = (dire + 1) % 4
			continue
		}
		if v == -2 {
			dire = (dire + 3) % 4
			continue
		}

		for i := 0; i < v; i++ {
			if hmap[[2]int{curx + directions[dire][0], cury + directions[dire][1]}] > 0 {
				break
			}
			curx += directions[dire][0]
			cury += directions[dire][1]
		}
		res = Max(res, curx*curx+cury*cury)
	}
	return res
}

/*
定义方向为 d { [0,1] 北 [1,0] 东，[0,-1] 南 [-1,0] 西  }
初始方向北 [0,1] 右转-1意味着方向 北->东->南->西 即 d[i++]
同理左转 d[i--] 防止取复数 d [ (i + 3)%4  ]
(curx,cury) 移动x步就是 (curx+d[i][0]*x,cury+d[i][0]*x  )
*/
func robotSim(commands []int, obstacles [][]int) int {
	dirs := [4][]int{
		[]int{0, 1},
		[]int{1, 0},
		[]int{0, -1},
		[]int{-1, 0},
	}
	direction := 0
	m := map[[2]int]int{}
	for _, v := range obstacles {
		m[[2]int{v[0], v[1]}] = 1
	}

	curX, curY := 0, 0
	d := 0
	for _, command := range commands {
		if command == -1 {
			direction = (direction + 1) % 4
			continue
		}
		if command == -2 {
			direction = (direction + 3) % 4
			continue
		}
		for i := 0; i < command; i++ {
			if m[[2]int{curX + dirs[direction][0], curY + dirs[direction][1]}] > 0 {
				break
			}
			curX += dirs[direction][0]
			curY += dirs[direction][1]
			d = Max(d, curX*curX+curY*curY)
		}

	}
	return d
}

func TestRobit(t *testing.T) {
	tests := []struct {
		name     string
		commands []int
		ob       [][]int
		wantRes  int
	}{
		{
			name:     "1",
			commands: []int{4, -1, 3},
			ob:       [][]int{},
			wantRes:  25,
		},
		{
			name:     "2",
			commands: []int{4, -1, 4, -2, 4},
			ob:       [][]int{[]int{2, 4}},
			wantRes:  65,
		},
		{
			name:     "3",
			commands: []int{6, -1, -1, 6},
			ob:       [][]int{},
			wantRes:  36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := RobotSim(tt.commands, tt.ob)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
