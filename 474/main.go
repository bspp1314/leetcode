package main

import "fmt"

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

func max(a, b int) int  {
	if a > b {
		return a
	}

	return b
}

func backpack(valuesAndWeight [2][]int,totalWeight int) int {
	dp := make([][]int,len(valuesAndWeight[0]))
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,totalWeight + 1 )
	}

	//放入第一个物品，填第一行列表
	for j := 0;j <len(valuesAndWeight[0]);j++{
		if valuesAndWeight[1][j] < totalWeight {
			dp[0][j] = valuesAndWeight[0][j]
		}
	}


	for i := 1 ; i < len(valuesAndWeight[0]) ; i++ {
		for j := totalWeight;j > 0;j-- {
			if j > valuesAndWeight[1][i]{
				dp[i][j] = max(dp[i-1][j],dp[i-1][j-valuesAndWeight[1][i]] + valuesAndWeight[0][i])
			}
		}
	}

	return dp[len(valuesAndWeight[0])-1][totalWeight]
}

func main() {

	out := backpack([2][]int{
		[]int{12,3,10,3,6},
		[]int{6,4,7,2,6},
	},15)
	fmt.Println(out)
	//in := []string{"10","0001","111001","1","0"}
	//out := findMaxForm(in, 3, 4)
	//fmt.Println(out)
}
