package grap

import "testing"

/*
review
https://leetcode.cn/problems/course-schedule-ii/
dfs遍历
*/
func FindOrder(numCourses int, prerequisites [][]int) []int {
	var (
		res     = []int{}
		visited = make([]int, numCourses)
		grap    = make([][]int, numCourses)
		valid   = true
		dfs     func(i int)
	)
	for _, v := range prerequisites {
		grap[v[0]] = append(grap[v[0]], v[1])
	}
	dfs = func(i int) {
		visited[i] = 1
		for _, v := range grap[i] {
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
		visited[i] = 2
		res = append(res, i)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
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
}
