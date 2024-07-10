/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var meow *ListNode = &ListNode{Next: head}
	head = meow
	var first, second *ListNode = head, head
	for first != nil && first.Next != nil {
		second, first = second.Next, first.Next.Next
	}
	second.Next = second.Next.Next
	return head.Next
}