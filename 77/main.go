package main

import "fmt"


// n k

func combine(n int,k int)[][]int  {
	var res [][]int

	var dfs func(sub []int,left int)


	dfs = func(sub []int, left int) {
		if (n -left +1)  < k  - len(sub){
			return
		}

		if len(sub) == k {
			temp := make([]int,len(sub))
			copy(temp,sub)
			res = append(res,temp)
			return
		}

		for i := left; i <= n ; i++ {
			newSub := append(sub,i)
			dfs(newSub,i+1)
		}
	}

	dfs([]int{},1)

	return res

}



func main() {
	fmt.Println( combine(5,2))
}
