package tree

import "testing"

/*
https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
数组元素不重复
还原树结构

	前序遍历                中序遍历

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *TreeNode {
	/*
		1. 中序遍历的第一个一定是根节点
		2. 根据根节点去找左右树的分界，当前函数只负责构建自己的节点
	*/
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for k, v := range inorder {
		if v == preorder[0] {
			i = k
			break
		}
	}
	root.Left = BuildTree(preorder[1:1+i], inorder[:i])
	root.Right = BuildTree(preorder[1+i:], inorder[i+1:])
	return root
}

/*
前序遍历 根节点 <左节点> <右节点>
中序遍历 <左节点> 根节点 <右节点>  这个前提是数组元素不重复
preorder[0]为根节点，根据这个去中序遍历找位置，可以获取 左右节点的数量
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func TestBuildTree(t *testing.T) {
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Log(inorderTraversal(root))
	t.Log(inorderTraversal(BuildTree([]int{-1}, []int{-1})))
}
