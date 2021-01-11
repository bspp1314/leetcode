package  main

import "fmt"


//var dp [][]bool

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	//dp = make([][]bool,len(s))
	//for i:=0;i<len(s);i++ {
	//	dp[i] = make([]bool,len(s))
	//	dp[i][i] = true
	//}
	//
	//for right := 0; right < len(s); right++ {
	//	for left := 0; left <=right ; left++ {
	//		if s[left] == s[right]  && (right -left <= 2 || dp[left+1][right-1]){
	//			dp[left][right] = true
	//		}
	//	}
	//}

	res := make([][]string,0)
	partitionHelp(s,0,len(s),[]string{},&res)

	return res
}

func isPalindrome(s string,left int,right int) bool  {
	for left < right {
		 if s[left] != s[right] {
		 	return false
		 }

		 left += 1
		 right -=1
	}

	return true
}


func partitionHelp(s string,left int,lenS int,sub []string,res *[][]string)  {
	if left == lenS {
		temp := make([]string,len(sub))
		copy(temp,sub)
		*res = append(*res,temp)
		return
	}

	//fmt.Println(sub)

	for i:=left;i < lenS;i++ {

		if !isPalindrome(s,left,i) {
			continue
		}
		newSub := append(sub,s[left:i+1])
		partitionHelp(s,i+1,lenS,newSub,res)
	}

}




func main() {

	fmt.Println(partition("bbabbaba"))
}
