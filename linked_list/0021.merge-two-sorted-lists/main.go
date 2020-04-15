package main

import "fmt"

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
type node struct {
	value    int
	nextNode *node
}
func main() {
	//head := new(ListNode)
	//head.Val = 3
	//ln2 := new(ListNode)
	//ln2.Val = 1
	//ln3 := new(ListNode)
	//ln3.Val = 4

	no1 := new(node)
	no2 := new(node)
	no3 := new(node)
	no4 := new(node)
	no5 := new(node)
	no1.value =4
	no2.value =2
	no3.value =4
	no4.value =1
	no5.value =7
	no1.nextNode = no2
	no2.nextNode = no3
	no3.nextNode = no4
	no4.nextNode = no5
	reverseNode := reverseNode(no1)
	printNode(reverseNode)



}
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	tmp := &ListNode{}
//	res := tmp
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			tmp.Next = l1
//			l1 = l1.Next
//		} else {
//			tmp.Next = l2
//			l2 = l2.Next
//		}
//		tmp = tmp.Next
//	}
//	if l1 == nil {
//		tmp.Next = l2
//	} else {
//		tmp.Next = l1
//	}
//	return res.Next
//}

/**
思路:
根据大小，插入链表。
 */
func reverseNode(head *node) *node {
	//  先声明两个变量
	//  前一个节点
	var preNode *node
	preNode = nil
	//  后一个节点
	nextNode := new(node)
	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.nextNode
		//  将头节点指向前一个节点
		head.nextNode = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	return preNode
}

func printNode(head *node) {
	for head != nil {
		//fmt.Print(head.value, "\t")
		fmt.Println(head)
		head = head.nextNode
	}
	fmt.Println()
}