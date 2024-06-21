package tree

import "testing"

/*
94 返回中序遍历  左子树->树根节点->右子树
590 返回前序遍历
n叉树的后续遍历
*/
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}
	res = append(res, root.Val)
	return res
}

func TestRoot(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	t.Log(inorderTraversal(root))
}
