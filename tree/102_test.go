package tree

func levelOrderTwo(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		temp := stack
		stack = nil
		val := []int{}
		for _, v := range temp {
			val = append(val, v.Val)
			if v.Left != nil {
				stack = append(stack, v.Left)
			}
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
		}
		res = append(res, val)
	}
	return res
}
