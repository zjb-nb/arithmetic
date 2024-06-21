package grap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
https://leetcode.cn/problems/course-schedule/
review
*/
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return true
}

/*
是否全部学完课程 == 图没有环结构
图用矩阵或者出边数组存储值
visted标记已经访问的节点
经过一个点置为1，未经过置为0  已经入栈为2，即该点和该点所有路径已经遍历且无环
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}
	res := []int{}
	visted := make([]int, numCourses)
	vaild := true
	var dfs func(u int)
	dfs = func(u int) {
		visted[u] = 1
		for _, v := range edges[u] {
			if visted[v] == 0 {
				dfs(v)
				if !vaild {
					return
				}
			} else if visted[v] == 1 {
				vaild = false
				return
			}
		}
		visted[u] = 2
		res = append(res, u)
	}
	for i := 0; i < numCourses && vaild; i++ {
		if visted[i] == 0 {
			dfs(i)
		}
	}
	return vaild
}

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		data    [][]int
		wantRes bool
	}{
		{
			name: "1",
			num:  2,
			data: [][]int{
				[]int{1, 0},
			},
			wantRes: true,
		},
		{
			name: "2",
			num:  2,
			data: [][]int{
				[]int{1, 0},
				[]int{0, 1},
			},
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CanFinish(tt.num, tt.data)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
