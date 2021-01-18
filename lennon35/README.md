---

---

### leetcode题库 35. 搜索插入位置

---
#### 原题信息

---
##### 原题链接:

https://leetcode-cn.com/problems/search-insert-position/

---
##### 难度等级
```
中等
```

---
##### 原题描述
```
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。
```

![example1](/img/lennon35/example.png)

---
#### 实现

---
##### 构思
```

```
---
##### 代码实现
```
func searchInsert(nums []int, target int) int {
	var i, n int
	for i, n = range nums {
		if target <= n {
			return i
		}
	}
	return i + 1
}
```
---
##### 代码链接

https://github.com/lennon-liu/leetcode/tree/main/lennon35

---
#### 测试结果

![lennon2](/img/lennon35/lennon35.png)

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
