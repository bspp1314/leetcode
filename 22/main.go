package main


func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := make([]string,0)
	var dfs func(sub []byte,left,right int)

	dfs = func(sub []byte,left,right int){
		if left == 0 && right == 0 {
			t := make([]byte,len(sub))
			copy(t,sub)
			res = append(res,string(t))
			return
		}

		if left == right {
			dfs(append(sub,'('),left-1,right)
		}else if left < right {
			if left > 0 {
				dfs(append(sub,'('),left-1,right)
			}

			dfs(append(sub,')'),left,right-1)
		}

	}

	dfs([]byte{},n,n)
	return res
}

func main() {

}
