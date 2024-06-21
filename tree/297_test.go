package tree

import (
	"strconv"
)

/*BNF*/
type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

func (c *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	res := ""
	res += "(" + c.Serialize(root.Left) + ")"
	res += strconv.Itoa(root.Val)
	res += "(" + c.Serialize(root.Right) + ")"
	return res
}

// ( X num X  ) num ( X num X )
func (c *Codec) Deserializes(data string) *TreeNode {
	var parse func() *TreeNode
	parse = func() *TreeNode {
		if data == "X" {
			return nil
		}
		node := &TreeNode{}
		data = data[1:]
		node.Left = parse()
		data = data[1:]
		i := 0
		for data[i] == '-' || data[i] >= '0' && data[i] <= '9' {
			i++
		}
		node.Val, _ = strconv.Atoi(data[:i])
		data = data[i+1:]
		node.Right = parse()
		data = data[1:]
		return node
	}
	return parse()
}

/*
树为空 X
树不为空 (<LEFT_SUB_TREE>)CUR_NUM(RIGHT_SUB_TREE)

巴科斯范式如下
T -> (T) num (T) | X
*/
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	left := "(" + this.serialize(root.Left) + ")"
	right := "(" + this.serialize(root.Right) + ")"
	return left + strconv.Itoa(root.Val) + right
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var parse func() *TreeNode
	parse = func() *TreeNode {
		if data[0] == 'X' {
			data = data[1:]
			return nil
		}
		node := &TreeNode{}
		data = data[1:]
		node.Left = parse()
		data = data[1:]
		i := 0
		for data[i] == '-' || '0' <= data[i] && '9' >= data[i] {
			i++
		}
		node.Val, _ = strconv.Atoi(data[:i])
		data = data[i+1:]
		node.Right = parse()
		data = data[1:]
		return node
	}
	return parse()
}
