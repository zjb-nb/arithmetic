package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
验证一棵树是否为二叉搜索树
有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
*/
func IsValidBST(root *TreeNode) bool {
	/*
		搜索树 左子树必须全部小于根节点，右子树反之 利用区间约束
		dfs(low,up,root)
		  root>low && root < up
	*/
	if root == nil {
		return true
	}
	var dfs func(left, right int, root *TreeNode) bool
	dfs = func(left, right int, root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val <= left || root.Val >= right {
			return false
		}
		l := dfs(left, root.Val, root.Left)
		r := dfs(root.Val, right, root.Right)
		return l && r
	}

	return dfs(math.MinInt, root.Val, root.Left) && dfs(root.Val, math.MaxInt, root.Right)
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

func TestBst(t *testing.T) {
	tests := []struct {
		name    string
		root    *TreeNode
		wantRes bool
	}{
		{
			name: "1",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			wantRes: true,
		},
		{
			name: "2",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			wantRes: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsValidBST(tt.root)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
