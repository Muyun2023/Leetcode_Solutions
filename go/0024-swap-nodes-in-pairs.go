func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    next := head.Next
    swapped := swapPairs(next.Next)
    
    next.Next, head.Next = head, swapped
    
    return next
}
