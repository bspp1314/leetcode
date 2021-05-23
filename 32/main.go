package main


func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}

	stack := make([]int,0)
	stack = append(stack, -1)
	res := 0

	for i:=0;i <len(s);i++ {
		if s[i] == '(' {
			stack = append(stack,i)
		}else{
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i - stack[len(stack)-1])
			}
		}
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
