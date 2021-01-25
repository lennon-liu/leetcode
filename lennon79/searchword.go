package main

import (
	"fmt"
)

func check(board [][]byte, word []byte, x, y, i int) bool {
	fmt.Println(x, y)
	if len(word) <= 0 || x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 {
		return false
	}
	if board[x][y] == word[0] {
		fmt.Println(board[x][y])
		fmt.Println(word[0])
		return true
	}
	return false
}

func exist1(board1 [][]byte, word []byte, x, y, i int, direction string) bool {
	board := board1

	if i == len(word)-1 {
		return true
	}
	if len(word) <= 0 || x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 {
		return false
	}
	if check(board, word[i:], x, y, i) {
		board[x][y] = '-'
		i++
	}
	if direction == "01" {
		exist1(board, word[i:], x+1, y, i, "10")
		exist1(board, word[i:], x, y+1, i, "01")
		exist1(board, word[i:], x-1, y, i, "-10")
	} else if direction == "10" {
		exist1(board, word[i:], x, y+1, i, "01")
		exist1(board, word[i:], x+1, y, i, "10")
		exist1(board, word[i:], x, y-1, i, "0-1")
	} else if direction == "0-1" {
		exist1(board, word[i:], x+1, y, i, "10")
		exist1(board, word[i:], x, y-1, i, "0-1")
		exist1(board, word[i:], x-1, y, i, "-10")
	} else if direction == "-10" {
		exist1(board, word[i:], x, y+1, i, "01")
		exist1(board, word[i:], x, y-1, i, "0-1")
		exist1(board, word[i:], x-1, y, i, "-10")
	}
	return false
}

func exist(board [][]byte, word string) bool {
	var x, y, i = 0, 0, 0
	words := []byte(word)
	return exist1(board, words, x, y, i, "01")
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "SEE"))
}
