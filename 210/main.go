package main
//统计课程安排图中每个节点的入度，生成 入度表 indegrees。
//借助一个队列 queue，将所有入度为 0 的节点入队。
//当 queue 非空时，依次将队首节点出队，在课程安排图中删除此节点 pre：
//并不是真正从邻接表中删除此节点 pre，而是将此节点对应所有邻接节点 cur 的入度 -1，即 indegrees[cur] -= 1。
//当入度 -1 后邻接节点 cur 的入度为 00，说明 cur 所有的前驱节点已经被 “删除”，此时将 cur 入队。



//func canFinish(numCourses int, prerequisites [][]int) bool {
//	var (
//		edges = make([][]int, numCourses)
//		indeg = make([]int, numCourses)
//		result []int
//	)
//
//	for _, info := range prerequisites {
//		edges[info[1]] = append(edges[info[1]], info[0])
//		indeg[info[0]]++
//	}
//
//	q := []int{}
//	for i := 0; i < numCourses; i++ {
//		if indeg[i] == 0 {
//			q = append(q, i)
//		}
//	}
//
//	for len(q) > 0 {
//		u := q[0]
//		q = q[1:]
//		result = append(result, u)
//		for _, v := range edges[u] {
//			indeg[v]--
//			if indeg[v] == 0 {
//				q = append(q, v)
//			}
//		}
//	}
//	return len(result) == numCourses
//}

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	//入度表
	edges := make([][]int, numCourses)
	indegree := make([]int,numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indegree[info[0]]++
	}


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


