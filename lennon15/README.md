### leetcode题库 15. 三数之和

---
#### 原题信息

---
##### 原题链接:

>https://leetcode-cn.com/problems/3sum/

---
##### 难度等级

>中等

---
##### 原题描述

>给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。


>示例 1：

>输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

>输入：nums = []
输出：[]
示例 3：

>输入：nums = [0]
输出：[]


>提示：
0 <= nums.length <= 3000
-105 <= nums[i] <= 105




#### 构思
>这个题比较容易超时，因为提交测试数组长度较长，在使用常规遍历后，加入多条件终止循环从而减少没有必要的遍历，办法上属于比较笨的办法。

---
#### 实现
---
##### 代码实现
```go
func threeSum(nums []int) [][]int {
	result:=[][]int{}
	sort.Ints(nums)
	temp:=1
	for index:=0;index<len(nums)-2;index+=1{
		if nums[index]>0{
			break
		}
		if temp==nums[index]{
			continue
		}else{
			temp=nums[index]
		}
		temp2:=100000
		for index2:=index+1;index2<len(nums)-1;index2+=1 {
			if nums[index]+nums[index2]>0{
				break
			}
			if nums[index]+nums[index2]+nums[len(nums)-1]<0{
				continue
			}
			if temp2==nums[index2]{
				continue
			}else{
				temp2=nums[index2]
			}
			temp3:=-1
			for index3:=len(nums)-1;index3>index2;index3-=1 {
				if nums[index]+nums[index2]+nums[index3]<0{
					break
				}
				if nums[index3]<0{
					break
				}
				if temp3==nums[index3]{
					continue
				}else{
					temp3=nums[index3]
				}
				if nums[index]+nums[index2]+nums[index3]==0{
					temp:=[]int{}
					temp=append(temp, nums[index],nums[index2],nums[index3])
					result=append(result, temp)
					break
				}
			}
		}
	}
	return result
}
```
---
##### 代码链接

>https://github.com/lennon-liu/leetcode/tree/main/lennon15

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
