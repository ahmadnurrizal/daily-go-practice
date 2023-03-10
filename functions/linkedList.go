package functions

type ListNode struct {
	Val  int
	Next *ListNode
}

// delete duplicate value in linked with ascending sort
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return slow.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
