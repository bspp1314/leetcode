package main

import "fmt"

func numTrees2(n int) int   {
	nMap := make([]int, n + 1)
	nMap[0], nMap[1] = 1, 1


	for i:= 2;i<= n;i++ {
		for j := 1;j <= i ;j++ {
			nMap[i] += nMap[j-1] * nMap[i-j]
		}

	}
	return nMap[n]
}
func numTrees(n int) int {
	nMap := make(map[int]int)
	nMap[0], nMap[1] = 1, 1
	return numTreesHelp(n,nMap)
}

func numTreesHelp(n int,nMap map[int]int) int  {
	res := 0
	for i:= 0;i< n ;i++ {
		left := i
		right := n - 1 - i

		if v,ok := nMap[left];ok {
			left = v
		} else{
			left = numTreesHelp(left,nMap)
		}

		if v,ok := nMap[right];ok {
			right = v
		}else{
			right = numTreesHelp(right,nMap)
		}

		res = res + left*right
	}

	nMap[n] = res

	return res
}
func main()  {
	out := numTrees(10)
	fmt.Println(out)
	fmt.Println(numTrees2(10))
}
