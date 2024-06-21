package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Array() []int {
	res := make([]int, 0, 5)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func MakeList(nums []int) *ListNode {
	var res *ListNode = new(ListNode)
	head := res
	for i := 0; i < len(nums); i++ {
		res.Next = new(ListNode)
		res = res.Next
		res.Val = nums[i]
	}
	return head.Next
}
