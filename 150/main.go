package main

import (
	"fmt"
	"strconv"
)

var tokenFuncs = map[string]func(a,b int )int {
	"*": func(a, b int) int {
		return a*b
	},
	"+": func(a, b int) int {
		return a+b
	},

	"-": func(a, b int) int {
		return a-b
	},

	"/": func(a, b int) int {
		return  a/b
	},
}
func evalRPN(tokens []string) int {
	if len(tokens) <= 2 {
		return 0
	}

	stack := make([]int,0)
	for i := 0; i < len(tokens); i++ {
		tokenFun,ok := tokenFuncs[tokens[i]]
		if ok {
			v1 := stack[len(stack)-1]
			v2 := stack[len(stack)-2]
			stack  = stack[0:len(stack)-2]
			stack = append(stack,tokenFun(v1,v2))
		} else{
			v,_ := strconv.Atoi(tokens[i])
			stack = append(stack,v)
		}

		fmt.Println(stack)
	}

	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
}
