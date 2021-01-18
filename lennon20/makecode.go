package lennon20

import "container/list"

func isValid(s string) bool {
	Stack := list.New()
	for _,code :=range(s){
		before := Stack.Back()
		if before==nil{
			Stack.PushBack(code)
			continue
		}
		temp := before.Value.(int32)
		if temp=='(' && string(code) ==")"{
			Stack.Remove(before)
			continue
		}
			Stack.PushBack(code)
	}
	if Stack.Len()>0{
		return(false)
	}else {
		return(true)
	}
}


func generateParenthesis(n int) []string {

	for num:=0;num <n;num++{
		for num1:=0;num1 <n;num1++{

		}
	}
	return []{}
}