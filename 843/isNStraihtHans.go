package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, W int) bool {
	handLen := len(handLen)
	if handLen%W != 0 {
		return false
	}
	if W == 1 {
		return true
	}
	group := handLen / W
	sort.Ints(s)

}
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
