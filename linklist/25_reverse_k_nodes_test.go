package linklist

import "testing"

//https://leetcode.cn/problems/reverse-nodes-in-k-group/ 12345 3 32145
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, tail := head, head
	for i := 0; i < k-1; i++ {
		if tail.Next == nil {
			return head
		}
		tail = tail.Next
	}
	next := tail.Next
	tail.Next = nil
	newNode := reverseList4(cur)
	cur.Next = reverseKGroup(next, k)
	return newNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	prev := new(ListNode)
	res := prev
	for head != nil {
		i := 0
		start := head
		for i < k-1 && head != nil {
			head = head.Next
			if head == nil {
				prev.Next = start
				return res.Next
			}
			i++
		}
		next := head.Next
		head.Next = nil
		prev.Next = swap(start)
		prev = start
		head = next
	}
	return res.Next
}

//递归
func reverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	start := head
	i := 0
	for i < k-1 && head != nil {
		head = head.Next
		if head == nil {
			return start
		}
		i++
	}

	next := head.Next
	head.Next = nil
	res := swap(start)
	start.Next = reverseKGroup1(next, k)
	return res
}

func TestGroup(t *testing.T) {
	head := MakeList([]int{1, 2, 3, 4, 5})
	t.Log(ReverseKGroup(head, 3).Array())
}
