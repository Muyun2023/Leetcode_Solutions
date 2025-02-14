
func pairSum(head *ListNode) int {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := reverse(slow.Next)
	// split into 2 lists
	slow.Next = nil

	first := head
	var maxSum int
	for first != nil && second != nil {
		maxSum = max(maxSum, first.Val+second.Val)
		first = first.Next
		second = second.Next
	}
	return maxSum
}

func reverse(node *ListNode) *ListNode {
	var curr, prev *ListNode = node, nil
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
