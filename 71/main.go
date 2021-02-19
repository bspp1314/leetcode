package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) == 0 {
		return ""
	}
	paths := strings.Split(path, "/")

	stack := make([]string, 0)
	for _, f := range paths {
		if f == "." || f == "" {
			continue
		} else if f == ".." {
			if len(stack) >= 1 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, f)
		}
	}
	return "/"+strings.Join(stack, "/")
}

func main() {
	// 0 1 2
	// 1 2 3
	// 3 4 5

	// 0 1 0
	// 1 2 3
	// 3 4 5
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
