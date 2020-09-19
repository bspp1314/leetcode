package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	dfs(&res,0,n,k,make([]int,0))
	return res
}

func dfs(res *[][]int,left int,n int,k int,q []int) {
	if k > (n-left) {
		return
	}

	if k == 1 {
		for i := left;i < n;i++ {
			tem := make([]int,len(q)+1)
			copy(tem,q)
			tem[len(tem)-1] = i+1
			*res = append(*res,tem)
		}

		return
	}

	for i:=left;i<n;i++{
		dfs(res,i+1,n,k-1,append(q,i+1))
	}
}

func main() {
	res := combine(4,2)
	fmt.Println(res)
}
