package main

import "fmt"

type Node struct {
	Next *Node
	Pre  *Node
	Key  int
	Val  int
}

type LRUCache struct {
	capacity int
	len      int
	head     *Node
	tail     *Node
	NodeMap  map[int]*Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		len:      0,
		head:     &Node{},
		tail:     &Node{},
		NodeMap:  make(map[int]*Node),
	}

	l.head.Key = -1
	l.head.Val = -1
	l.tail.Key = -1
	l.tail.Val = -1

	l.head.Next = l.tail
	l.tail.Pre = l.head

	return l
}

func (this *LRUCache) Get(key int) int {
	node := this.get(key)
	if node == nil {
		return -1
	}

	return node.Val
}

func (this *LRUCache) get(key int) *Node {
	node, exist := this.NodeMap[key]
	if !exist {
		return nil
	}

	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre


	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node

	return node
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}

	node := this.get(key)
	if node != nil {
		node.Val = value
		return
	}

	node = &Node{Val: value, Key: key}
	this.len++
	this.NodeMap[key] = node

	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node

	if this.len > this.capacity {
		v := this.tail.Pre.Key
		this.tail.Pre = this.tail.Pre.Pre
		this.tail.Pre.Next = this.tail
		delete(this.NodeMap,v)
	}
}

func Print(a *Node) {
	v := a
	for v != nil {
		fmt.Printf(" [%d,%d] ", v.Key, v.Val)
		v = v.Next
	}

	fmt.Println()
}

func main() {
	lru := Constructor(3)
	lru.Put(1, 2)
	lru.Put(2, 10)
	Print(lru.head)
	lru.Put(3, 30)
	Print(lru.head)
	lru.Put(4, 40)
	Print(lru.head)
	lru.get(4)
	Print(lru.head)
	//lru.get(2)
	//Print(lru.head)
}
