// @Title  leetcode题库 2.两数相加
// @Description  leetcode题库 2.两数相加 代码实现（golang）
// @Author  lennon
// @Update  2021-01-18

package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listNode := new(ListNode)
	temp := listNode
	if l1 != nil {
		temp.Val += l1.Val
	}
	if l2 != nil {
		temp.Val += l2.Val
	}

	for {
		tnextnode := new(ListNode)
		if temp.Val > 9 {
			temp.Val = temp.Val - 10
			tnextnode.Val = 1
			temp.Next = tnextnode
			temp = tnextnode
		}
		if l1.Next == nil && l2.Next == nil {
			temp.Next = nil
			break
		} else {
			temp.Next = tnextnode
			temp = tnextnode
		}
		if l1.Next != nil {
			l1 = l1.Next
			temp.Val += l1.Val
		}
		if l2.Next != nil {
			l2 = l2.Next
			temp.Val += l2.Val
		}
	}
	return listNode
}

//这是一条测试语句
func main() {
	listNode1 := new(ListNode)
	list1 := []int{2, 4, 3}
	listNode1.Val = list1[0]
	tepp := listNode1
	for i := 1; i < len(list1); i += 1 {
		temp := new(ListNode)
		tepp.Next = temp
		tepp = temp
		tepp.Val = list1[i]
	}
	listNode2 := new(ListNode)
	list2 := []int{5, 6, 4}
	listNode2.Val = list2[0]
	tepp2 := listNode2
	for i := 1; i < len(list2); i += 1 {
		temp := new(ListNode)
		tepp2.Next = temp
		tepp2 = temp
		tepp2.Val = list2[i]
	}
	fmt.Println(listNode1)
	fmt.Println(listNode2)
	result := addTwoNumbers(listNode1, listNode2)
	fmt.Println(result)
}
