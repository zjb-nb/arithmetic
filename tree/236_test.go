package tree

/*
二叉树的最近公共祖先
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return nil

}

/*
分类讨论
1.当前节点是空节点，返回当前节点Nil
2.当前节点是p/q，返回当前节点
3. 其他：
      1.左右子树都找到了，返回当前节点
			2.只有左子树找到，返回左子树结构
			3.只有右子树找到，返回右子树结构
			4.都没找到，返回空节点nil

*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return right
}
