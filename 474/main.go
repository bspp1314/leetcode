package main

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int,m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int,n+1)
	}

	for i := 0; i < len(strs); i++ {
		nums :=getOneNum(strs[i])

		for zeroes := m;zeroes >= nums[0];zeroes-- {
			for ones :=n;ones >= nums[1];ones-- {
				max := func(a,b int) int {
					if a > b {
						return a
					}else{
						return b
					}
				}
				dp[zeroes][ones] = max(1 + dp[zeroes - nums[0]][ones - nums[1]],
					dp[zeroes][ones])
			}
		}

	}

	return dp[m][n]
}

func getOneNum(s string) []int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			n++
		}
	}

	return []int{n,len(s)-n}
}

func main() {
	//in := []string{"10","0001","111001","1","0"}
	//out := findMaxForm(in, 3, 4)
	//fmt.Println(out)
}
