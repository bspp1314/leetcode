package main

import "fmt"

var byteToInt = map[byte]int{
	'1':1,
	'2':2,
	'3':3,
	'4':4,
	'5':5,
	'6':6,
	'7':7,
	'8':8,
	'9':9,
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	v := countAndSay(n-1)

	byteMap := make(map[byte]int)
	res := ""
	for i:=0;i<len(v);i++{
		if i == 0 {
			byteMap[v[i]] = 1
		}else{
			if byteMap[v[i]] == 0 {
				k := byteMap[v[i-1]]
				if k != 0 {
					res += fmt.Sprintf("%d%d",k,byteToInt[v[i-1]])
					delete(byteMap,v[i-1])
				}
			}
			byteMap[v[i]]++
		}
	}


	k := byteMap[v[len(v)-1]]
	res += fmt.Sprintf("%d%d",k,byteToInt[v[len(v)-1]])

	return res
}


func main() {
	// 1 2 2 1
	// 1 1 2 2 1 1
	fmt.Println(countAndSay(10))

}
