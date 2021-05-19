package main

import "fmt"

var keysMap = map[byte]string{
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"rqrs",
	'8':"tuv",
	'9':"wxyx",
}
func letterCombinations(digits string) []string {
	res := make([]string,0)

	var dfs func(begin int,sub []byte)


	dfs = func(begin int, sub []byte) {
		if begin == len(digits) {
			tem := make([]byte,len(sub))
			copy(tem,sub)
			res = append(res,string(tem))
			return
		}

		keys := keysMap[digits[begin]]
		for i:=0;i <len(keys);i++ {
			newSub := append(sub,keys[i])
			dfs(begin+1,newSub)
		}
	}

	dfs(0,[]byte{})

	return res
}

func main() {
	out := letterCombinations("")
	fmt.Println(out)
}