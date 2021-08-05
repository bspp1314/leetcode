package main

import "fmt"

func canFinish2(numCourses int,prerequisites [][]int)  {
	//标记学习 i 课程之前需要学习的课程
	edges := make([][]int, numCourses)


	for i := 0; i <len(prerequisites) ; i++ {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]],prerequisites[i][0])
	}

	result := make([]int, 0)
	visited := make([]int, numCourses)
	valid := true


	var dfs  func(u int)

	// 验证课程
	dfs = func(i int) {
		if visited[i] == 2 {
			return
		}

		visited[i] = 1 //标记课程被访问过
		for _,v := range edges[i] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			}else if visited[v] == 1 {
				valid = false
			}
		}

		visited[i] = 2
		result = append(result,i)
	}
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

}
