package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
type ListNode struct{
	Next *ListNode
	Val int
}
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{Next: nil,
		Val: 0}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := NewListNode()
	tmpHead := head
	flag := 0
	pl, pr := l1, l2
	if pl == nil {
		return pr
	}
	if pr == nil {
		return pl
	}
	for {
		if flag != 0 || pl != nil || pr != nil {
			newNode := NewListNode()
			if pl != nil {
				newNode.Val += pl.Val
				pl = pl.Next
			}
			if pr != nil {
				newNode.Val += pr.Val
				pr = pr.Next
			}
			if flag == 1 {
				newNode.Val += flag
			}
			flag = newNode.Val / 10
			newNode.Val = newNode.Val % 10
			tmpHead.Next = newNode
			tmpHead = tmpHead.Next
		} else {
			return head.Next
		}
	}
}
func main() {
	l1 := NewListNode()
	l1.Val = 2
	l2 := NewListNode()
	l2.Val = 4
	l3 := NewListNode()
	l3.Val = 3
	l1.Next = l2
	l2.Next = l3

	r1 := NewListNode()
	r1.Val = 5
	r2 := NewListNode()
	r2.Val = 6
	r3 := NewListNode()
	r3.Val = 4
	r1.Next = r2
	r2.Next = r3

	ret := addTwoNumbers(l1, r1)
	for {
		if ret != nil {
			fmt.Println(ret.Val)
			ret = ret.Next
		} else {
			break
		}
	}
}
