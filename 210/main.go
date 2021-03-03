package main
//统计课程安排图中每个节点的入度，生成 入度表 indegrees。
//借助一个队列 queue，将所有入度为 00 的节点入队。
//当 queue 非空时，依次将队首节点出队，在课程安排图中删除此节点 pre：
//并不是真正从邻接表中删除此节点 pre，而是将此节点对应所有邻接节点 cur 的入度 -1−1，即 indegrees[cur] -= 1。
//当入度 -1−1后邻接节点 cur 的入度为 00，说明 cur 所有的前驱节点已经被 “删除”，此时将 cur 入队。

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	//入度表

}
func findOrder(numCourses int, prerequisites [][]int) []int {
	//领接矩阵
	edges := make([][]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	visited := make([]int, numCourses)

	result := make([]int, 0)

	valid := true

	var dfs func(u int)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}


		visited[u] = 2
		result = append(result, u)
	}




	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}

	reverse(0,len(result)-1,result)


	return result

}

func reverse(left,right int,arr []int)   {
	for left < right{
		arr[left],arr[right] = arr[right],arr[left]
		left++
		right--
	}
}


