package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}



func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var dfs  func(node *Node) *Node

	vis := make(map[*Node]*Node)

	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		copyNode,ok := vis[node]
		if ok {
			return copyNode
		}



		copyNode = &Node{
			Val:node.Val,
		}

		vis[node] = copyNode

		copyNode.Next = dfs(node.Next)
		copyNode.Random = dfs(node.Random)

		return copyNode
	}

	return dfs(head)
}

func copyRandomList3(head *Node) *Node {
	if head == nil {
		return nil
	}



	vis := make(map[*Node]*Node)
	vis[head] = &Node{Val:head.Val}
	q := []*Node{head}

	for len(q) != 0 {
		n := q[0]
		q = q[1:len(q)]



		if n.Next != nil {
			_,ok := vis[n.Next]
			if !ok {
				cn := &Node{Val:n.Next.Val}
				vis[n.Next]= cn
				q = append(q,n.Next)
			}
		}
		vis[q[0]].Next = vis[n.Next]

		if n.Random != nil {
			_,ok := vis[n.Random]
			if !ok {
				cn := &Node{Val:n.Random.Val}
				vis[n.Random]= cn
				q = append(q,n.Random)
			}
		}

		vis[q[0]].Random= vis[n.Random]


	}


	return vis[head]
}

func main() {
}