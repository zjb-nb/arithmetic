package tree

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = InvertTree(root.Right)
	root.Right = InvertTree(root.Left)
	return root
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
