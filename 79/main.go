package main

import "fmt"

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j:=0;j<len(board[0]);j++ {
			if search(board,word,i,j,0){
				return true
			}
		}
	}

	return false
}


func  search(board [][]byte, word string,i,j,k int)  bool{
	if k >= len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[k] {
		return false
	}

	c := board[i][j]
	board[i][j] += ','

	result := search(board, word, i - 1, j, k + 1) ||
		search(board, word, i + 1, j, k + 1) ||
		search(board, word, i, j - 1, k + 1) ||
		search(board, word, i, j + 1, k + 1)

	board[i][j] = c

	return result
}




func main() {
	//in := [][]byte{
	//	{'A', 'B', 'C', 'E'},
	//	{'S', 'F', 'C', 'S'},
	//	{'A', 'D', 'E', 'E'},
	//}
	in := [][]byte{
		{'A','B'},
	}

	fmt.Println(exist(in,"BA"))

}
