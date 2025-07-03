package pkg

type Node struct {
	key, value int
	pre, next  *Node
}
type LRUCache struct {
	capacity int           // 容量大小
	cache    map[int]*Node // 缓存 map
	head     *Node         // 哨兵头节点
	tail     *Node         // 哨兵尾节点
}

type function interface {
	Get(key int) int
	Put(key int, value int)
	MoveToHead(node *Node)
	AddToHead(node *Node)
	RemoveNode(node *Node)
	RemoveTail() *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		// 已存在，返回节点值
		l.MoveToHead(node)
		return node.value
	} else {
		// 不存在，返回 -1
		return -1
	}
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		// 已存在，更新值到头节点
		node.value = value
		l.MoveToHead(node)
	} else {
		// 不存在，则创建节点
		newNode := &Node{key: key, value: value}
		l.cache[key] = newNode
		l.AddToHead(newNode)
		// 如果超出容量就删除最久未使用的节点
		if len(l.cache) > l.capacity {
			tailNode := l.RemoveTail()
			delete(l.cache, tailNode.key)
		}
	}
}

func (l *LRUCache) RemoveNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (l *LRUCache) MoveToHead(node *Node) {
	l.RemoveNode(node)
	l.AddToHead(node)
}

func (l *LRUCache) AddToHead(node *Node) {
	node.pre = l.head
	node.next = l.head.next
	l.head.next.pre = node
	l.head.next = node
}

func (l *LRUCache) RemoveTail() *Node {
	node := l.tail.pre
	l.RemoveNode(node)
	return node
}
