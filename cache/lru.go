package cache

import "sync"

var _ cache = (*LRUCache)(nil)

// double linked list
type Node struct {
	key  interface{}
	val  interface{}
	pre  *Node
	next *Node
}

type LRUCache struct {
	sync.Mutex

	capacity int
	item     map[interface{}]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		item:     make(map[interface{}]*Node),
	}
}

func (cache *LRUCache) setHead(node *Node) {
	cache.Lock()
	defer cache.Unlock()
	node.next = cache.head
	node.pre = nil
	if cache.head != nil {
		cache.head.pre = node
	}
	cache.head = node
	if cache.tail == nil {
		cache.tail = node
	}
}

func (cache *LRUCache) delete(node *Node) {
	cache.Lock()
	cache.Unlock()
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		// means node is the head of cache
		cache.head = node.next
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		cache.tail = node.pre
	}
}

func (cache *LRUCache) Set(key, val interface{}) error {
	if v, ok := cache.item[key]; ok {
		v.val = val
		cache.delete(v)
		cache.setHead(v)
		return nil
	}

	node := &Node{
		key: key,
		val: val,
	}
	if len(cache.item) >= cache.capacity {
		cache.delete(cache.tail)
		delete(cache.item, cache.tail.key)
	}
	cache.setHead(node)
	cache.Lock()
	cache.item[key] = node
	cache.Unlock()

	return nil
}

func (cache *LRUCache) Get(key interface{}) (interface{}, error) {
	if v, ok := cache.item[key]; ok {
		cache.delete(v)
		cache.setHead(v)
		return v.val, nil
	}
	return nil, nil
}
