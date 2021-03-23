package main

type Node struct {
	next     *Node
	pre      *Node
	key      int
	value    int
	frequent int
}

type ListNode struct {
	head *Node
	tail *Node
	size int
}

// 双向链表辅助函数：添加一个节点到头节点后
func (this *ListNode) addNode(node *Node) {
	head := this.head
	node.next = head.next
	node.next.pre = node
	node.pre = head
	head.next = node
}

// 双向链表辅助函数，删除一个节点
func (this *ListNode) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}



type LFUCache struct {
	capacity    int
	size        int
	minFrequent int
	frequentMap map[int]*ListNode
	cacheMap    map[int]*Node
}

func NewLru(cap int) *LFUCache {
	l := Constructor(cap)
	return &l
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:    capacity,
		size:        0,
		minFrequent: 0,
		cacheMap:    make(map[int]*Node),
		frequentMap: make(map[int]*ListNode),
	}
}


// LFUCache辅助函数：将节点从对应的频次双向链表中删除
func (this *LFUCache) remove(node *Node) {
	this.frequentMap[node.frequent].removeNode(node)
	this.frequentMap[node.frequent].size--
}

// LFUCache辅助函数：将节点添加进对应的频次双向链表，没有则创建
func (this *LFUCache) add(node *Node) {
	if listNode, exist := this.frequentMap[node.frequent]; exist {
		listNode.addNode(node)
		listNode.size++
	} else {
		listNode = &ListNode{&Node{}, &Node{}, 0}
		listNode.head.next = listNode.tail
		listNode.tail.pre = listNode.head
		listNode.addNode(node)
		listNode.size++
		this.frequentMap[node.frequent] = listNode
	}
}
// LFUCache辅助函数：移除一个key
func (this *LFUCache) evictNode() {
	listNode := this.frequentMap[this.minFrequent]
	delete(this.cacheMap, listNode.tail.pre.key)
	listNode.removeNode(listNode.tail.pre)
	listNode.size--
}

// LFUCache辅助函数：获取一个key和修改一个key都会增加对应key的访问频次，可以独立为一个方法，完成如下任务：
// 1. 将对应node从频次列表中移出
// 2. 维护minFrequent
// 3. 该节点访问频次++，移动进下一个访问频次链表
func (this *LFUCache) triggerVisit(node *Node) {
	this.remove(node)
	if node.frequent == this.minFrequent && this.frequentMap[node.frequent].size == 0 {
		this.minFrequent++
	}
	node.frequent++
	this.add(node)
}

func (this *LFUCache) Get(key int) int {
	if node, exist := this.cacheMap[key]; exist {
		this.triggerVisit(node)
		return node.value
	}
	return -1
}


func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if node, exist := this.cacheMap[key]; exist {
		this.triggerVisit(node)
		this.cacheMap[key].value = value
	} else {
		newNode := &Node{
			next:     nil,
			pre:      nil,
			key:      key,
			value:    value,
			frequent: 1,
		}
		if this.size < this.capacity {
			this.add(newNode)
			this.size++
			this.minFrequent = 1
		} else {
			this.evictNode()
			this.add(newNode)
			this.minFrequent = 1
		}
		this.cacheMap[key] = newNode
	}
}




func main() {

}
