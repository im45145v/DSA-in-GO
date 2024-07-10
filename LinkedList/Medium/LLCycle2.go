/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	var count bool
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			count = true
			break
		}
	}
	if count != true {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}