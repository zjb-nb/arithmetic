package grap

import "testing"

/*
review
https://leetcode.cn/problems/redundant-connection/
输入: edges = [[1,2], [1,3], [2,3]]
输出: [2,3]
判断是否有多余的环
*/
func FindRedundantConnection(edges [][]int) []int {
	return []int{}
}

/*
在添加边的时候判断是否形成环

添加出边数组的时候，两个方向都要添加

每次添加后没法确定，之前无环的地方不会出现环，因此标记已入栈没有意义
*/
func findRedundantConnection(edges [][]int) []int {
	n := 0
	for _, v := range edges {
		n = Max(n, v[0])
		n = Max(n, v[1])
	}

	grap := make([][]int, n+1)
	visted := make([]int, n+1)
	var dfs func(u, father int)
	hasCycle := false

	dfs = func(u, father int) {
		visted[u] = 1
		for _, v := range grap[u] {
			if v == father {
				continue
			}
			if visted[v] == 1 {
				hasCycle = true
				return
			}
			dfs(v, u)
			if hasCycle {
				return
			}
		}
		visted[u] = 0
	}

	for _, v := range edges {
		grap[v[0]] = append(grap[v[0]], v[1])
		grap[v[1]] = append(grap[v[1]], v[0])

		for i := 1; i <= n; i++ {
			dfs(i, 0)
			if hasCycle {
				return v
			}
		}
	}

	return []int{}
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestFindConnection(t *testing.T) {
	t.Log(findRedundantConnection([][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}))
}
