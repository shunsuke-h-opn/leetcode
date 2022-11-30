package main

import "fmt"

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func middleNode(head *ListNode) *ListNode {

// }

func reverseList(head *ListNode) *ListNode {
	var list *ListNode
	for head != nil {
		list = &ListNode{Val: head.Val, Next: list}
		head = head.Next
	}
	return list
}

func reverseList2(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func lengthNode(n *ListNode) (result int) {
	count := 0
	temp := n
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
		count += 1
	}
	return count
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	len := lengthNode(head)

	middle := len / 2

	index := 0
	var returnVal *ListNode
	fmt.Println(len)
	fmt.Println(middle)
	fmt.Println(head)
	for index < len {

		if index >= middle {

			returnVal = head
			break
		}
		head = head.Next
		fmt.Println(head)
		fmt.Println(returnVal)
		index++
	}
	return returnVal
}

func main() {
	fmt.Println("-----------start leetcode1481--------------")

	c := &ListNode{3, nil}
	b := &ListNode{2, c}
	a := &ListNode{1, b}
	list := a

	// fmt.Println(list)
	// fmt.Println(list.Val)
	// fmt.Println(list.Next)
	// fmt.Println(list.Next)
	// fmt.Println(list.Next.Next.Val)
	// fmt.Println(list.recursiveLength())

	middleVal := middleNode(list)
	fmt.Println("-----------line 98--------------")
	fmt.Println(middleVal)

	// fmt.Println(val)
	// fmt.Println(reverseList(list))

	fmt.Println("-----------end leetcode1481--------------")
}
