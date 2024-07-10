/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		later := curr.Next
		curr.Next = prev
		prev = curr
		curr = later
	}
	return prev
}
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var slow, fast, one *ListNode
	slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	two := reverseList(slow.Next)
	one = head
	for two != nil {
		if one.Val != two.Val {
			return false
		}
		one = one.Next
		two = two.Next
	}
	return true

}