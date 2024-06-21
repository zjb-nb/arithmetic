package tree

import (
	"math"
)

/*
验证一棵树是否为二叉搜索树
有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func IsValidBST(root *TreeNode) bool {
	return true
}

/*
1. 左子树值 必须全部小于 node 同理右子树 利用区间约束
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(low, up int, root *TreeNode) bool
	dfs = func(low, up int, root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= up {
			return false
		}
		left_vaild := dfs(low, root.Val, root.Left)
		right_vaild := dfs(root.Val, up, root.Right)
		if !left_vaild || !right_vaild {
			return false
		}

		return true
	}
	vaild := dfs(math.MinInt, math.MaxInt, root)
	return vaild
}
