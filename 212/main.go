package main

import "fmt"

type Trie struct {
	next   map[byte]*Trie
	word   string
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	for i := 0; i < len(word); i++ {
		if t.next[word[i]] == nil {
			t.next[word[i]] = &Trie{
				next:   make(map[byte]*Trie),
			}
		}

		t = t.next[word[i]]
	}

	t.word = word
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	trie := &Trie{
		next: make(map[byte]*Trie),
		word: "",
	}

	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	res := make([]string,0)

	var backtracking func(row,col int,parent *Trie)

	backtracking = func(row, col int, parent *Trie) {
		letter := board[row][col]
		currNode := parent.next[letter]

		if currNode != nil {
			if currNode.word != "" {
				res = append(res,currNode.word)
			}
			currNode.word = ""
		}

		board[row][col] = '#'


		x := []int{1,0,-1,0}
		y := []int{0,1,0,-1}

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				newRow := row + x[i]
				newCol := col + y[i]

				if newRow < 0 || newRow >= len(board) || newCol < 0 || newCol >= len(board[0]) {
					continue
				}else{
					if currNode.next[board[newRow][newCol]] != nil {
						backtracking(newRow,newCol,currNode)
					}
				}

			}
		}

		board[row][col] = letter
		if len(currNode.next) == 0 {
			delete(parent.next,letter)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if trie.next[board[i][j]] != nil {
				backtracking(i,j,trie)
			}
		}
	}

	return res

}
func main() {
	fmt.Println(findWords([][]byte{
		{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'},
	},[]string{"oath","pea","eat","rain"}))
}
