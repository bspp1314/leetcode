package main

import (
	"fmt"
)

func isValidSerialization(preorder string) bool {
	// n := len(preorder)
	// stack := []int{1}
	//
	//for i := 0; i < n; {
	//	if len(stack) == 0 {
	//		return false
	//	}
	//
	//	if preorder[i] == ',' {
	//		i++
	//		continue
	//	}
	//
	//	if preorder[i] == '#' {
	//		stack[len(stack)-1]--
	//		if stack[len(stack)-1] == 0 {
	//			stack = stack[:len(stack)-1]
	//		}
	//		i++
	//	}else{
	//		for i <n  && preorder[i] != ',' {
	//			i++
	//		}
	//
	//		stack[len(stack)-1]--
	//		if stack[len(stack)-1] == 0 {
	//			stack = stack[:len(stack)-1]
	//		}
	//		stack = append(stack, 2)
	//
	//	}
	//
	//}
}


func main()  {
	in := "3,4,#,#,1,#,#"
	fmt.Println(isValidSerialization(in))
}
