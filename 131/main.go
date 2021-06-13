package  main

import "fmt"


//var dp [][]bool
func partition(s string) [][]string {
	l := len(s)
	if l == 0 {
		return [][]string{}
	}

	f := make([][]int8, l)
	for i := range f {
		f[i] = make([]int8, l)
	}

	// 0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
		return f[i][j]
	}


	res := make([][]string,0)

	var dfs func(left int,sub []string)

	dfs = func(left int, sub []string) {
		if left == l {
			temp := make([]string,len(sub))
			copy(temp,sub)
			res = append(res,temp)
			return
		}

		for i := left; i < l; i++ {
			if isPalindrome(left,i) > 0 {
				newSub := append(sub,s[left:i+1])
				dfs(i+1,newSub)
			}

		}
	}

	dfs(0,[]string{})
	return res
}


func partition3(s string) [][]string {
	l := len(s)
	if l == 0 {
		return [][]string{}
	}

	dp := make([][]bool,l)
	for i := range dp {
		dp[i] = make([]bool, l)
		for j := range dp[i] {
			dp[i][j] = true
		}

		fmt.Println(dp[i])
	}

	for left := l-1;left >= 0;left-- {
		for right := left+1;right < l;right++ {
			if left == 4 && right == 5 {
				fmt.Println( s[left] == s[right])
				fmt.Println(dp[left+1][right-1])
				fmt.Printf("dp[%d][%d] = %t \n",left+1,right-1,dp[left+1][right-1])
			}
			dp[left][right] = s[left] == s[right] && dp[left+1][right-1]
		}
	}

	fmt.Println(dp[4][5])

	res := make([][]string,0)

	var dfs func(left int,sub []string)

	dfs = func(left int, sub []string) {
		if left == l {
			temp := make([]string,len(sub))
			copy(temp,sub)
			res = append(res,temp)
			return
		}

		for i := left; i < l; i++ {
			if dp[left][i] {
				newSub := append(sub,s[left:i+1])
				dfs(i+1,newSub)
			}

		}
	}

	dfs(0,[]string{})
	return res
}


func partitionK(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	fmt.Println(f[4][5])

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}



func main() {

	fmt.Println(partition3("bbabbb"))
	//fmt.Println(partitionK("bbabbb"))

}
