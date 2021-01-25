---
title: 28. 实现 strStr()
cover: img/leetcode.png
tags: 
	- 简单难度
	- golang实现
	- 字符串
categories: 
	- leetcode
---

### leetcode题库 28. 实现 strStr()

---
#### 原题信息

---
##### 原题链接:

> https://leetcode-cn.com/problems/implement-strstr/
>

---
##### 难度等级
> 简单

---
##### 原题描述
> 实现 strStr() 函数。
> 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。



>示例 1:
>输入: haystack = "hello", needle = "ll"
>输出: 2

>######示例 2:
>输入: haystack = "aaaaa", needle = "bba"
>输出: -1


>说明:
>当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。




---
#### 构思
> 字符串整行匹配

---
#### 实现
---
##### 代码实现
```go
func strStr(haystack string, needle string) int {
	i:=0
	for{
		if i> len(haystack) || len(haystack[i:])< len(needle){
			return -1
		}
		if haystack[i:i+len(needle)]==needle {
			return i
		}
		i+=1
	}
}
```
---
##### 代码链接

https://github.com/lennon-liu/leetcode/tree/main/lennon28

---
#### 测试结果

![lennon28](/img/lennon28/lennon28.png)

----
#### 优化与总结
```

```

---
> 来源：力扣（LeetCode）
> 链接：https://leetcode-cn.com/problems/add-two-numbers
> 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。