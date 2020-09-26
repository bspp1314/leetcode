package  main

import "fmt"

func partition(s string) [][]string {
	fmt.Println(s)
	if len(s) == 0 {
		return [][]string{}
	}

	if len(s) == 1 {
		return [][]string{[]string{s}}
	}




	res := make([][]string,0)
	for i:=0;i<len(s)-1;i++ {
		res1 := partition(s[:i+1])
		res2 := partition(s[i+1:])

		if len(res1) == 0 {
			res = append(res,res1...)
			continue
		}
		//
		if len(res2) == 0 {
			res = append(res,res2...)
			continue
		}

		for j:=0;j<len(res1);j++ {
			for m := 0; m < len(res2); m++ {
				res = append(res,append(res1[j],res2[m]...))
			}
		}
	}

	if havePalindrome(s) {
		res = append(res,[]string{s})
	}

	return  res
}



func havePalindrome(s string) bool {
	left := 0
	right := len(s) -1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left ++
		right--
	}

	return true
}

func main() {
	s := "aaaa"
	for i := 0; i < len(s)-1; i++ {
		fmt.Println(s[:i+1])
		fmt.Println(s[i+1:])
	}

	out := partition(s)
	fmt.Println(out)


}
