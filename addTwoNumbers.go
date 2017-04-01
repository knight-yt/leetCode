package main

/**
* leetCode addTwoNumbers
* @author 79597374@qq.com
 */
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	val := l1.Val + l2.Val
	step := 0
	if val >= 10 {
		val = val % 10
		step = 1
	}
	l3 := &ListNode{Val: val}
	n2 := l2.Next
	n1 := l1.Next

	n := l3
	for {
		if n1 == nil && n2 == nil {
			break
		}

		val1 := 0
		if n1 != nil {
			val1 = n1.Val
		}

		val2 := 0
		if n2 != nil {
			val2 = n2.Val
		}

		val = val1 + val2 + step
		if val >= 10 {
			val = val % 10
			step = 1
		} else {
			step = 0
		}
		n.Next = &ListNode{Val: val}
		if n2 != nil {
			n2 = n2.Next
		}
		if n1 != nil {
			n1 = n1.Next
		}
		n = n.Next
	}
	if step == 1 {
		n.Next = &ListNode{Val: 1}
	}
	return l3
}

func Insert(l *ListNode, n *ListNode) bool {
	if l.Next == nil {
		l.Next = n
		return true
	}

	head := l
	for head.Next != nil {
		if head.Next.Next == nil {
			head.Next.Next = n
			return true
		}
		head = head.Next
	}
	return true
}

func main() {
	i := []int{8}
	j := []int{6}

	l1 := NewListNode(1)
	l2 := NewListNode(0)
	for _, val := range i {
		n := NewListNode(val)
		Insert(l1, n)
	}

	for _, val := range j {
		n := NewListNode(val)
		Insert(l2, n)
	}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3.Val)
	for l3.Next != nil {
		fmt.Println(l3.Next.Val)
		l3.Next = l3.Next.Next
	}
}
