package main

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	return IsValidSerialization(strings.Split(preorder,","))
}

func IsValidSerialization(preorder []string) bool {
	if len(preorder) == 1 {
		return true
	}

	if preorder[0] == "#" {
		return false
	}

	for i:=1;i<len(preorder);i++ {
		res := true
		if len(preorder) == 7 && i == 3 {
			fmt.Println(preorder[1:i+1])
			fmt.Println(preorder[i+1:])
		}
		res = res && IsValidSerialization(preorder[1:i+1])
		if i != len(preorder) -1 {
			res = res && IsValidSerialization(preorder[i+1:])
		}

		if res {
			return true
		}
	}

	return false





	return res
}

func main()  {
	in := "3,4,#,#,1,#,#"
	fmt.Println(isValidSerialization(in))
}
