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

func combine2(n int, k int) [][]int {
	var res [][]int
	dfs2([]int{},0,n,k,&res)
	return res
}

func dfs2(subNumbs []int,left,n,k int,res *[][]int) {
	if (n - left) < k - len(subNumbs) {
		return
	}

	if len(subNumbs) == k {
		temp := make([]int,len(subNumbs))
		copy(temp,subNumbs)
		*res = append(*res,temp)
		return
	}

	for i := left; i < n; i++ {
		dfs2(append(subNumbs,i+1),i+1,n,k,res)
	}

}

func main() {
	fmt.Println( combine(5,2))
	fmt.Println( combine2(5,2))
}
