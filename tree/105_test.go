package tree

/*
数组元素不重复
还原树结构

	前序遍历                中序遍历

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
*/
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	node.Left = BuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	node.Right = BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return node
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
