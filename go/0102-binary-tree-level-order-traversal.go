func levelOrder(root *TreeNode) [][]int {
		if root == nil {
			return nil
		}
		
		var result [][]int
		level := 0
		q := make([]*TreeNode,1)
		q[0] = root
		for len(q) > 0 {
			curLen := len(q)
			
			for i := 0; i < curLen; i++ {
			
				node := q[0]
				q = q[1:]
				
				if len(result) <= level {
					result = append(result,[]int{node.Val})
				}else {
					result[level] = append(result[level],node.Val)
				}
				
				if node.Left != nil {
					q = append(q,node.Left)
				}
				if node.Right != nil{
					q = append(q,node.Right)
				}
			}
			level++
		}
		return result
		
	}
