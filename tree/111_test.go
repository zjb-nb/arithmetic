package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil && root.Left != nil {
		return minDepth(root.Left) + 1
	}
	return Min(minDepth(root.Left)+1, minDepth(root.Right)+1)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
