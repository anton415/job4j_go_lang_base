package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{
		size: size,
	}
}

func (l *LruCache) Put(key string, value string) {
	node := &Node{Key: key, Value: value, Prev: l.Head}
	if l.Head != nil {
		l.Head.Next = node
	} else {
		l.Tail = node
	}
	l.Head = node

	count := 1
	for node := l.Head; node != nil; node = node.Prev {
		if count == l.size {
			l.Tail = node
			node.Prev = nil
			return
		}
		count++
	}
}

func (l *LruCache) Get(key string) *string {
	node := l.find(key)
	if node == nil {
		return nil
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
		if node.Prev != nil {
			node.Prev.Next = node.Next
		} else {
			l.Tail = node.Next
		}
		node.Prev = l.Head
		node.Next = nil
		l.Head.Next = node
		l.Head = node
	}
	return &node.Value
}

func (l *LruCache) find(key string) *Node {
	for node := l.Head; node != nil; node = node.Prev {
		if node.Key == key {
			return node
		}
	}
	return nil
}
