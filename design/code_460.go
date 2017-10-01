package design

// https://leetcode.com/problems/lfu-cache
type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Score int
}

type LFUCache struct {
	cache    map[int]int
	keyScore map[int]*Node
	capacity int
	tail     *Node
}

func Constructor(capacity int) LFUCache {
	lc := LFUCache{
		cache:    make(map[int]int),
		keyScore: make(map[int]*Node),
		capacity: capacity,
	}
	return lc
}

func (lc *LFUCache) Get(key int) int {
	v, found := lc.cache[key]
	if !found {
		return -1
	}
	lc.hitKey(key, v)

	return v
}

func (lc *LFUCache) Put(key int, value int) {
	if lc.capacity == 0 {
		return
	}
	if len(lc.cache) >= lc.capacity {
		if _, found := lc.cache[key]; !found {
			lc.removeMinScore()
		}
	}
	lc.hitKey(key, value)
}

func (lc *LFUCache) hitKey(key int, value int) {
	r := lc.keyScore[key]
	if r == nil {
		r = &Node{
			Prev: lc.tail,
			Key:  key,
		}
		lc.keyScore[key] = r
		lc.tail = r
	}

	lc.cache[key] = value
	r.Score += 1
	for r.Prev != nil && r.Score >= r.Prev.Score {
		lc.moveForward(r)
	}
}

func (lc *LFUCache) removeMinScore() {
	tail := lc.tail
	if tail == nil {
		return
	}
	delete(lc.cache, tail.Key)
	delete(lc.keyScore, tail.Key)

	tail = tail.Prev
	if tail != nil && tail.Prev != nil {
		tail.Prev.Next = nil
	}
	lc.tail = tail
}

func (lc *LFUCache) moveForward(node *Node) {
	prev := node.Prev
	node.Prev = prev.Prev
	if prev.Prev != nil {
		prev.Prev.Next = node
	}

	if node.Next != nil {
		node.Next.Prev = prev
	}

	prev.Next = node.Next
	prev.Prev = node
	node.Next = prev

	if lc.tail == node {
		lc.tail = prev
	}
}
