package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
二叉树的最近公共祖先
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*
		寻找两个节点的公共祖先，也就是返回左右子树都存在pq的第一个节点
		1.如果当前节点是p/q，则直接返回
		2. 不是的话分别讨论
		   - nil则直接返回nil
			 - 左右子树都没找到nil
			 - 左/右找到返回左/右，都找到返回当前节点
	*/
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
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
	return left
}

func TestLow(t *testing.T) {
	root := BuildTree([]int{3, 5, 6, 2, 7, 4, 1, 0, 8}, []int{6, 5, 7, 2, 4, 3, 0, 1, 8})
	// t.Log(preorderTraversal(root))
	tests := []struct {
		name    string
		p       int
		q       int
		wantRes int
	}{
		{
			name:    "1",
			q:       5,
			p:       1,
			wantRes: 3,
		},
		{
			name:    "2",
			q:       5,
			p:       4,
			wantRes: 5,
		},
		{
			name:    "3",
			q:       6,
			p:       8,
			wantRes: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LowestCommonAncestor(root, &TreeNode{Val: tt.p}, &TreeNode{Val: tt.q})
			require.NotNil(t, res)
			assert.Equal(t, tt.wantRes, res.Val)
		})
	}
}
