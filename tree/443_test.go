package tree

/*
1.BFS
*/
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	m := map[string]struct{}{}
	for _, v := range bank {
		m[v] = struct{}{}
	}
	if _, ok := m[endGene]; !ok {
		return -1
	}
	stack := []string{startGene}
	for step := 0; stack != nil; step++ {
		temp := stack
		stack = nil
		for _, cur := range temp {
			for k, v := range cur {
				for _, y := range "ACTG" {
					if v != y {
						nxt := cur[:k] + string(y) + cur[k+1:]
						if _, ok := m[nxt]; ok {
							if nxt == endGene {
								return step + 1
							}
							delete(m, nxt)
							stack = append(stack, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}

/*
2.预处理
*/
func diffOne(s, t string) (diff bool) {
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func minMutation1(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	m := len(bank)
	adj := make([][]int, m)
	endIndx := -1
	for i, s := range bank {
		if s == endGene {
			endIndx = i
		}
		for j := i + 1; j < m; j++ {
			if diffOne(s, bank[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	if endIndx == -1 {
		return -1
	}

	var q []int
	vis := make([]bool, m)
	for i, s := range bank {
		if diffOne(startGene, s) {
			q = append(q, i)
			vis[i] = true
		}
	}
	for step := 1; q != nil; step++ {
		temp := q
		q = nil
		for _, cur := range temp {
			if cur == endIndx {
				return step
			}
			for _, nxt := range adj[cur] {
				if !vis[nxt] {
					vis[nxt] = true
					q = append(q, nxt)
				}
			}
		}
	}
	return -1
}
