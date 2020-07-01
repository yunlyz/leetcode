package golang

type ListNode struct {
    Val  int
    Next *ListNode
}

// addTwoNumbers 两数相加
// 解题方法：1.遍历链表
// 解题思路：遍历两个链表，对应的索引值相加。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    pre := &ListNode{}
    current := pre
    carry, p, q := 0, l1, l2

    for p != nil || q != nil {
        sum, x, y := 0, 0, 0
        if p != nil {
            x = p.Val
        }
        if q != nil {
            y = q.Val
        }
        sum = x + y + carry
        carry = sum / 10
        sum = sum % 10

        current.Next = &ListNode{Val: sum}
        current = current.Next
        if p != nil {
            p = p.Next
        }
        if q != nil {
            q = q.Next
        }
    }

    if carry != 0 {
        current.Next = &ListNode{Val: carry}
    }
    return pre.Next
}
