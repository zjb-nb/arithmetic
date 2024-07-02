package grap

import "testing"

/*
review
https://leetcode.cn/problems/course-schedule-ii/
*/
func FindOrder(numCourses int, prerequisites [][]int) []int {
	/*
		返回学完所有课程的安排顺序，即返回图的遍历过程
		1.图节点的状态分为未搜索0，搜索中1，已搜索2
		2.当图搜索过程中发现节点为1则表示有环，不可能学完
	*/
	res := make([]int, 0, numCourses)
	edges := make([][]int, numCourses)
	for _, v := range prerequisites {
		edges[v[0]] = append(edges[v[0]], v[1])
	}
	flag := true
	tmp := make([]int, numCourses)
	var dfs func(i int)

	dfs = func(i int) {
		tmp[i] = 1
		for _, v := range edges[i] {
			if tmp[v] == 0 {
				dfs(v)
				if !flag {
					return
				}
			} else if tmp[v] == 1 {
				flag = false
				return
			}
		}
		tmp[i] = 2
		res = append(res, i)
	}

	for k, v := range tmp {
		if v == 0 {
			dfs(k)
		}
		if !flag {
			return []int{}
		}
	}
	return res
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[0]] = append(edges[info[0]], info[1])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}

	return result
}

func TestFindCours(t *testing.T) {
	t.Log(FindOrder(3, [][]int{[]int{1, 0}}))
	t.Log(FindOrder(4, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}}))
}
