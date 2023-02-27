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
