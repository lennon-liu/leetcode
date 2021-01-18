---

---

### leetcode题库 2.两数相加

---
#### 原题信息

---
##### 原题链接:

https://leetcode-cn.com/problems/add-two-numbers/

---
##### 难度等级
```
中等
```

---
##### 原题描述
```
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
```

![example1](/img/lennon2/example1.png)
```
提示：
每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零
```
---
#### 实现

---
##### 构思
```

```
---
##### 代码实现
```
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
```
---
##### 代码链接

https://github.com/lennon-liu/leetcode/tree/main/lennon2

---
#### 测试结果

![lennon2](/img/lennon2/lennon2.png)

----
#### 优化与总结
```

```

---
```
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
```
