func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }
    
    for len(lists) > 1 {
        var mergedLists []*ListNode
        for i := 0; i < len(lists); i += 2 {
            l1 := lists[i]
            var l2 *ListNode
            if (i + 1) < len(lists) {
                l2 = lists[i + 1]
            }
            mergedLists = append(mergedLists, mergeList(l1, l2))
        }
        lists = mergedLists
    }
    return lists[0]
}


