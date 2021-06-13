package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}


	visit := make(map[*Node]*Node)

	var dfs  func(root *Node) *Node

	dfs = func(root *Node) *Node {
		if root == nil {
			return nil
		}

		copyNode,ok := visit[root]
		if ok {
			return copyNode
		}

		copyNode = &Node{
			Val:       root.Val,
			Neighbors: make([]*Node,0),
		}
		
		visit[root] = copyNode

		for i := 0; i < len(root.Neighbors); i++ {
			copyNode.Neighbors = append(copyNode.Neighbors,dfs(root.Neighbors[i]))
		}

		return copyNode
	}

	return dfs(node)
}

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}

	if len(node.Neighbors) == 0 {
		return &Node{Val: node.Val}
	}

	visited := make(map[*Node]*Node)

	// 将题目给定的节点添加到队列
	levelNodes := []*Node{node}
	// 克隆第一个节点并存储到哈希表中
	visited[node] = &Node{node.Val, []*Node{}}

	for len(levelNodes) != 0 {
		n := levelNodes[0]
		levelNodes = levelNodes[1:]

		for i := 0; i < len(n.Neighbors); i++ {
			if _, ok := visited[n.Neighbors[i]]; !ok {
				// 如果没有被访问过，就克隆并存储在哈希表中
				visited[n.Neighbors[i]] = &Node{n.Neighbors[i].Val, []*Node{}}
				// 将邻居节点加入队列中
				levelNodes = append(levelNodes, n.Neighbors[i])
			}

			// 更新当前节点的邻居列表
			visited[n].Neighbors = append(visited[n].Neighbors, visited[n.Neighbors[i]])

		}
	}

	return visited[node]

}


