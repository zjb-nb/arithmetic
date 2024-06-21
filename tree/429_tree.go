package tree

/*
review
n叉树的dfs
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
*/
func LevelOrder(root *Node) [][]int {
	res := [][]int{}
	return res
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for stack != nil {
		q := stack
		level := []int{}
		stack = nil
		for _, v := range q {
			level = append(level, v.Val)
			stack = append(stack, v.Children...)
		}
		res = append(res, level)
	}
	return res
}
