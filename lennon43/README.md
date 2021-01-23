---
title: 43. 字符串相乘
cover: img/leetcode.png
tags: 
	- 中等难度
	- golang实现
	- 数组
categories: 
	- leetcode
---

### leetcode题库 43. 字符串相乘

---
#### 原题信息

---
##### 原题链接:

> https://leetcode-cn.com/problems/multiply-strings/
>

---
##### 难度等级
> 中等

---
##### 原题描述
> 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

![example1](/img/lennon43/example.png)

---
#### 实现

---
##### 构思
> 

---
##### 代码实现
```go
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var result [110 + 110]int
	length := (len(num1) - 1) + (len(num2) - 1) + 1
	for i1 := len(num1) - 1; i1 >= 0; i1-- {
		for i2 := len(num2) - 1; i2 >= 0; i2-- {
			numx := (int(num1[i1]) - 48) * (int(num2[i2]) - 48)
			if numx >= 10 {
				result[i1+i2+1] += numx % 10
				result[i1+i2] += numx / 10
			} else {
				result[i1+i2+1] += numx
			}
			for i3 := i1 + i2 + 1; i3 >= 0; i3-- {
				temp := result[i3]
				if temp >= 10 {
					result[i3] = temp % 10
					result[i3-1] += temp / 10
				} else {
					break
				}

			}
		}
	}
	str := ""
	for i3 := 0; i3 <= length; i3++ {
		str += string(result[i3] + 48)
	}
	if str[0] == 48 {
		str = str[1:]
	}
	return str
}
```
---
##### 代码链接

https://github.com/lennon-liu/leetcode/tree/main/lennon43

---
#### 测试结果

![lennon2](/img/lennon43/lennon43.png)

----
#### 优化与总结
```

```

---
> 来源：力扣（LeetCode）
> 链接：https://leetcode-cn.com/problems/add-two-numbers
> 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。