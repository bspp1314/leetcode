package main

import (
	"fmt"
)

func Min(a,b int) int   {
	if a < b {
		return a
	}

	return b
}

func Max(a,b int) int   {
	if a > b {
		return a
	}

	return b
}
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	maxSide := 0

	for i:=0;i< rows;i++ {
		for j := 0; j < columns; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			pMaxSide := Min(columns-j,rows-i)

			for k := 1;k < pMaxSide;i++ {
				//对角为 0，
				flag := true
				if matrix[i+k][j+k] == '0' {
					break
				}

				for m := 0;m < k;m ++ {
					if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
						flag = false
						break

					}
				}

				if flag {
					maxSide = Max(maxSide,k+1)
				}else{
					break
				}

			}

		}
	}

	return maxSide

}

func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	res := 0
	dp := make([][]int,len(matrix))

	for i:=0;i<len(matrix);i++ {
		dp[i] = make([]int,len(matrix[i]))
		for j:=0;j<len(matrix[i]);j++ {
			if i == 0 || j == 0  {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				}

				res = Max(dp[i][j],res)
				continue
			}

			if matrix[i][j] == '1' {
				dp[i][j] = Min(dp[i][j-1],Min(dp[i-1][j],dp[i-1][i-1])) + 1
				res = Max(dp[i][j],res)
			}
		}
	}



	return res * res

}

func main() {
	o := maximalSquare2([][]byte{
		[]byte{'1','0','1'},
		[]byte{'1','1'},
	})

	fmt.Println(o)
}
