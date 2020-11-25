package main

import (
	"fmt"
)

func isStraihtHand(hand []int, begin int, W int) b {
	end := make([]int, W)
	j := 0
	for i := begin; i < len(hand); i++ {
		if begin == 0 {
			end[0] = hand[0]
			j++
			continue
		}

		if end[j] == hand[i]-1 {
			j++
			end[j] = hand[i]
		} else {
			for end[j] == hand[i] {
				i++
			}
		}

	}
}
func main() {
	fmt.Println("vim-go")
}
