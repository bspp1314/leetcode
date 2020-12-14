package main



// 深度优先遍历
func findCircleNum(M [][]int) int {
	visited := make([]bool,len(M))

	queue := make([]int,0)
	res := 0

	for i:=0;i<len(M);i++ {
		if visited[i] {
			continue
		}
		queue  = append(queue,i)

		for len(queue) != 0 {
			s := queue[0]
			queue = queue[1:]

			// 访问点s相邻的节点
			for j:=0;j < len(M);j++ {
				if M[s][j] == 1 && !visited[j] {
					queue = append(queue,j)
				}
			}
		}

		res++
	}

	return res
}
