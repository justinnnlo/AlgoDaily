package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// naive approach
// pu the nodes into an array and reversively put the items into the result
// time		O(2n)
// space	O(n)
// but it beats 100% LOL
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := []*ListNode{head}
	cur := head
	for cur.Next != nil {
		list = append(list, cur.Next)
		cur = cur.Next
	}
	dump := &ListNode{0, nil}
	dumpCur := dump
	for i := len(list) - 1; i >= 0; i-- {
		dumpCur.Next = &ListNode{list[i].Val, nil}
		dumpCur = dumpCur.Next
	}
	return dump.Next
}

// 2nd approach: inplace replacement(learned from others)
// https://leetcode.com/explore/learn/card/linked-list/219/classic-problems/1204/
// 1->2->3
// 2->1->3
// tine		O(n)
// space 	O(1) since in-place
// beats 100%
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for head.Next != nil { // since we are replacing the head.next, the loop should end when there is no next
		temp := head.Next     // the middle node
		head.Next = temp.Next // or head.Next.Next
		temp.Next = cur       // put the middle node in front of cur
		cur = temp            // assign cur as head
	}
	return cur
}

// helper
func arr2list(arr []int) *ListNode {
	dump := &ListNode{0, nil}
	dumpCur := dump
	for i := 0; i < len(arr); i++ {
		dumpCur.Next = &ListNode{arr[i], nil}
		dumpCur = dumpCur.Next
	}
	return dump.Next
}

func printList(l *ListNode) {
	c := l
	for c != nil {
		fmt.Println(c.Val)
		c = c.Next
	}
}

func main() {
	a := arr2list([]int{1, 2, 3, 4, 5})
	b := reverseList1(a)
	printList(b)
}