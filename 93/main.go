package main

import (
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	res := make([]string,0)
	dfs([]string{},s,&res,0)
	return res
}


func dfs(ip []string,s string,res *[]string,left int)  {
	if left == len(s) && len(ip) == 4  {
		*res = append(*res,strings.Join(ip,"."))
		return
	}

	//剩余字符过长
	if len(s) - left  >  (4 - len(ip)) * 3 {
		return
	}

	// 剩余字符过短
	if len(s) -left < 4 - len(ip) {
		return
	}

	v := 0
	for i := left; i < left +3 && i < len(s); i++ {
		v = v * 10 + int(s[i] - '0')
		if v < 0 || v > 255 {
			continue
		}

		ip = append(ip,s[left:i+1])
		dfs(ip,s,res,i+1)
		ip = ip[0:len(ip)-1]

		if v == 0 {
			break
		}
	}
}

func main() {
	fmt.Println(restoreIpAddresses("1111"))
}