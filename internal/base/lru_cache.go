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

}

func (l *LruCache) Get(key string) *string {
	return nil
}

func (l *LruCache) find(key string) *Node {
	for node := l.Head; node != nil; node = node.Next {
		if node.Key == key {
			return node
		}
	}
	return nil
}
