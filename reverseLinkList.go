/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		newHead := head.Next
		head.Next.Next = head
		head.Next = nil
		return newHead
	}

	first, second, third := head, head.Next, head.Next.Next

	first.Next = nil
	for first != second {
		second.Next = first
		first = second
		second = third
		if third.Next != nil {
			third = third.Next
		}
	}

	return first

}