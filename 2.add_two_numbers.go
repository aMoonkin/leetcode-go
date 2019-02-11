/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	current := &head

	p := l1
	q := l2

	carry := 0

	for {
		v1 := 0
		v2 := 0

		if p != nil {
			v1 = p.Val
			p = p.Next
		}

		if q != nil {
			v2 = q.Val
			q = q.Next
		}

		sum := carry + v1 + v2
		carry = sum / 10
		sum = sum % 10

		current.Next = &ListNode{sum, nil}
		current = current.Next

		if q == nil && p == nil {
			break
		}
	}
	if carry != 0 {
		current.Next = &ListNode{carry, nil}
	}
	return head.Next
}