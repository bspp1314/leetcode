package main

import "fmt"


func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	v := countAndSay(n-1)
	//1 1
	//2 11
	//3 21
	//4 1211
	//5 111221

	res := ""
	c := v[0]
	num := 1

	for i:=1;i < len(v);i++ {
		if v[i] == c {
			num++
		}else{
			res+=fmt.Sprintf("%d%c",num,c)
			c = v[i]
			num = 1
		}
	}

	res+=fmt.Sprintf("%d%c",num,c)

	return res

}


func main() {
	// 1 2 2 1
	// 1 1 2 2 1 1
	fmt.Println(countAndSay(10))

}
