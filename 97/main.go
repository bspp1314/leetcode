package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	n := len(s1)
	m := len(s2)
	k := len(s3)
	if m+n != k {
		return false
	}

	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true

	// 定义 f[i][j] 表示 Si 的前 i 个元素和 Sj 前 j 个元素
	// 素是否能交错组成s3的前 i+j 个元素。
	// 那么 s3[i+j-1] == s[i-1] && f【i-1][j]
	// 要么 s3[i+j-1] == s[j-1] && f [i][j-1]

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i-1][j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}

		}
	}

	return f[n][m]
}

func main() {
	v := isInterleave("aabcc", "dbbca", "aadbbcbcac")
	fmt.Println(v)
}
