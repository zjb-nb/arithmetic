package tree

func MaxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil {
		return MaxDepth(root.Right) + 1
	}
	return Max(MaxDepth(root.Left)+1, MaxDepth(root.Right)+1)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return maxDepth(root.Right) + 1
	}
	if root.Right == nil {
		return maxDepth(root.Left) + 1
	}
	return Max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
