package main

import "fmt"

//numCourses = 2, prerequisites = [[1,0],[0,1]]
//{{0,10},{3,18},{5,5},{6,11},{11,14},{13,1},{15,1},{17,4}}

// need false,

func canFinish2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	queue := []int{}
	visited := make([]int, numCourses)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < len(edges); i++ {
		fmt.Println(edges[i])
	}

	queue = []int{prerequisites[0][1]}

	for len(queue) != 0  {
		qNode := queue[0]
		queue = queue[1:len(queue)]
		visited[qNode] = 1
		for i := 0; i < len(edges[qNode]) ; i++ {
			if visited[edges[qNode][i]] == 0 {
				queue = append(queue,edges[qNode][i])
			}else if visited[edges[qNode][i]] == 1 {
				return false
			}
		}

		//visited[qNode] = 2
	}

	return true



}
func canFinish(numCourses int, prerequisites [][]int) bool {
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
		visited[u] = 1 //标记其被访问过
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

	return valid
}

func main() {
	fmt.Println(canFinish(3,[][]int{{1,0},{0,2},{2,1}}))
	fmt.Println(canFinish(20,[][]int{{0,10},{3,18},{5,5},{6,11},{11,14},{13,1},{15,1},{17,4}}))

	fmt.Println(canFinish2(3,[][]int{{1,0},{0,2},{2,1}}))
	fmt.Println(canFinish2(20,[][]int{{0,10},{3,18},{5,5},{6,11},{11,14},{13,1},{15,1},{17,4}}))
}
