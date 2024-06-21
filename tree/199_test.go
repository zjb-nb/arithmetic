package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		n := len(stack)
		temp := stack
		stack = nil
		for k, v := range temp {
			if k == n-1 {
				res = append(res, v.Val)
			}
			if v.Left != nil {
				stack = append(stack, v.Left)
			}
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
		}
	}
	return res
}
