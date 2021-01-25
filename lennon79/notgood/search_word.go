//
package main

import "fmt"

func check_in_list(clist [][2]int, c0 int, c1 int) bool {
	for _, listx := range clist {
		if listx[0] == c0 && listx[1] == c1 {
			return false
		}
	}
	return true
}

func search(board [][]byte, w byte, x int, y int, clist [][2]int) (int, int) {
	for i1, list1 := range board {
		for i2, ww := range list1 {
			if x == i1 && ((i2-y < -1) || (i2-y > 1)) {
				continue
			}
			if y == i2 && ((i1-x < -1) || (i1-x > 1)) {
				continue
			}
			if y > 0 {
				if y != i2 && x != i1 {
					continue
				}
			}
			if ww == w {
				if check_in_list(clist, i1, i2) {
					return i1, i2
				}
			}
		}
	}
	return -1, -1
}

func exist(board [][]byte, word string) bool {
	//fmt.Println(board)
	index1, index2 := -1, -1
	listx1 := [2]int{}
	listx2 := [][2]int{}
	for _, w := range word {
		index1, index2 = search(board, byte(w), index1, index2, listx2)
		listx1[0] = index1
		listx1[1] = index2
		listx2 = append(listx2, listx1)
		fmt.Println(index1, index2)
		if index1 < 0 && index2 < 0 {
			return false
		}
	}
	return true
}

//func main(){
//	board:=[][]byte{
//		{'A','B'},
//	{'C','D'},
//	}
//	fmt.Println(exist(board,"ABCD"))
//}
func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "SEE"))
}
