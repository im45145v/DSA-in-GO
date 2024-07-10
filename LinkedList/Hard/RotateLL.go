/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	var dummy = &ListNode{Next: head}
	count, node := 0, head
	for node.Next != nil {
		count++
		node = node.Next
	}
	count++
	j := k % count
	node.Next = head
	node = head
	for i := 1; i < count-j; i++ {
		node = node.Next
	}
	dummy.Next = node.Next
	node.Next = nil
	return dummy.Next
}