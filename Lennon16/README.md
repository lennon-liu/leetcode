---
title: 16. 最接近的三数之和
cover: img/leetcode.png
tags: 
	- 简单难度
	- golang实现
	- 字符串
categories: 
	- leetcode
---

### leetcode题库 16. 最接近的三数之和

---
#### 原题信息

---
##### 原题链接:

>https://leetcode-cn.com/problems/3sum-closest/

---
##### 难度等级

>中等

---
##### 原题描述

>给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
>返回这三个数的和。假定每组输入只存在唯一答案。



>示例 1：
>输入：nums = [-1,2,1,-4], target = 1
 输出：2
 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


>提示：
3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4




#### 构思
>首先对字符串排序可以在合适的位置减少没有必要的循环，遍历前两个数，在寻找第三个数，第三个数可以常规遍历，加入限制条件也可以


---
#### 实现
---
##### 代码实现
```go
func twodivde(nums []int, target int)int{
	//fmt.Println(nums)
	if len(nums)==2{
		if math.Abs(float64(nums[0]-target))-math.Abs(float64(nums[1]-target))>0{
			return nums[1]
		}else{
			return nums[0]
		}
	}
	length:= len(nums)-1
	if length%2==0{
		length=length/2
	}else{
		length=(length-1)/2
	}
	if target == nums[length]{
		return target
	}else if target > nums[length]{
		return twodivde(nums[length:],target)
	}else{
		return twodivde(nums[:length+1],target)
	}
}

func threeSumClosest(nums []int, target int) int {
	result:=0
	lastx:=float64(10000)

	ttt:=0
	sort.Ints(nums)
	for i1 := 0;i1<=len(nums)-3;i1++{
		lastnum3:=9999
		for i2:=i1+1;i2<=len(nums)-2;i2++{
			num3:=(target-nums[i1])-nums[i2]
			if num3>nums[len(nums)-1]{
				ttt= nums[len(nums)-1]+nums[i2]+nums[i1]
			}else if num3<nums[i2+1]{
				ttt= nums[i2+1]+nums[i2]+nums[i1]
			}else{
				ttt=twodivde(nums,num3)+nums[i2]+nums[i1]
			}
			if ttt-target==0{
				return ttt
			}
			if math.Abs(float64(ttt-target))-math.Abs(lastx)<0{
				lastx=float64(ttt-target)
				result=ttt
			}
			if ttt-target>0 && ttt-lastnum3>0{
				break
			}else{
				lastnum3=ttt
			}
		}
	}
	return result
}
```
---
##### 代码链接

>https://github.com/lennon-liu/leetcode/tree/main/lennon16

---
#### 测试结果

![lennon2](/img/lennon15/lennon15.png)

----
#### 优化与总结
```

```

---

>来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
