/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//	func hasCycle(head *ListNode) bool {
//	    slow,fast:=head,head
//	    for fast!=nil && fast.Next!=nil {
//	        slow=slow.Next
//	        fast=fast.Next.Next
//	        if slow==fast{return true}
//	    }
//	    return false
//	}
func hasCycle(head *ListNode) bool {
	visitedNodes := make(map[*ListNode]bool)
	curr := head
	for curr != nil {
		if visitedNodes[curr] != false {
			return true
		}
		visitedNodes[curr] = true
		curr = curr.Next
	}
	return false
}