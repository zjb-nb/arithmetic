package main

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
		res.Val = nums[i]
		if i == len(nums)-1 {
			break
		}
		res.Next = new(ListNode)
		res = res.Next
	}
	return head
}
