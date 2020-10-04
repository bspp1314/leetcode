package main

import "fmt"

func main() {
	depth := 4
	for i:=0;i< depth;i++ {
		for j := 0; j < 1 << i; j++ {
			fmt.Printf("o ")
		}
		fmt.Println()
	}
}
