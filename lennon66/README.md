---
title: leetcode题库 66. 加一
cover: img/leetcode.png
tags: 
	- 简单难度
	- golang实现
	- 数组
categories: 
	- leetcode
---
### leetcode题库 66. 加一

---
#### 原题信息

---
##### 原题链接:

> https://leetcode-cn.com/problems/search-insert-position/
>

---
##### 难度等级
> 中等

---
##### 原题描述
> 给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
>
> 最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
>
> 你可以假设除了整数 0 之外，这个整数不会以零开头。
>

![example1](/img/lennon66/example.png)

> 提示：
> 1 <= digits.length <= 100
> 0 <= digits[i] <= 9

---
#### 实现

---
##### 构思
> 

---
##### 代码实现
```go
func plusOne(digits []int) []int {
	digits[len(digits)-1] = digits[len(digits)-1] + 1
	max := len(digits) - 1
	for {
		if digits[max] == 10 {
			if max == 0 {
				digits[max] = 0
				digits = append([]int{1}, digits...)
			} else {
				digits[max] = 0
				digits[max-1] = digits[max-1] + 1
			}
		} else {
			break
		}
		if max < 1 {
			break
		}
		max -= 1
	}
	return digits
}
```
---
##### 代码链接

> https://github.com/lennon-liu/leetcode/tree/main/lennon35
>

---
#### 测试结果

![lennon2](/img/lennon66/lennon66.png)

----
#### 优化与总结
> 

---
> 来源：力扣（LeetCode）
> 链接：https://leetcode-cn.com/problems/add-two-numbers
> 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。