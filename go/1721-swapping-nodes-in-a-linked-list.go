func swapNodes(head *ListNode, k int) *ListNode {
	curr := head
	for i := 0; i < k-1; i++ {
		curr = curr.Next
	}

	left := curr
	right := head

	for curr.Next != nil {
		curr = curr.Next
		right = right.Next
	}

	left.Val, right.Val = right.Val, left.Val

	return head
}
