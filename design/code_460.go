// Package design
// https://leetcode.com/problems/lfu-cache
package design

import "fmt"

type Node struct {
	Prev  *Node
	Next  *Node
	Key   int
	Score int
}

func (node *Node) moveForward() {
	if node.Prev == nil {
		return
	}

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
}

type LFUCache struct {
	cache    map[int]int
	keyRank  map[int]*Node
	capacity int
	rankTail *Node
}

func Constructor(capacity int) LFUCache {
	lc := LFUCache{
		cache:    make(map[int]int),
		keyRank:  make(map[int]*Node),
		capacity: capacity,
	}
	return lc
}

func (lc *LFUCache) Get(key int) int {
	v, ok := lc.cache[key]
	if !ok {
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
		lc.delTail()
	}
	lc.hitKey(key, value)

}

func (lc *LFUCache) hitKey(key int, value int) {
	r := lc.keyRank[key]
	if r == nil {
		lc.cache[key] = value
		r = &Node{
			Prev: lc.rankTail,
			Key:  key,
			Next: nil,
		}
		lc.rankTail = r
	}

	r.Score += 1
	if len(lc.cache) >= 2 {
		fmt.Println("move before")
		fmt.Println("head", lc.rankTail.Prev.Key, lc.rankTail.Prev.Score)
		fmt.Println("tail", lc.rankTail.Key, lc.rankTail.Score)
		fmt.Println("----")
	}

	for r.Prev != nil && r.Score > r.Prev.Score {
		r.moveForward()
	}

	if len(lc.cache) >= 2 {
		fmt.Println("move after")
		fmt.Println("head", lc.rankTail.Prev.Key, lc.rankTail.Prev.Score)
		fmt.Println("tail", lc.rankTail.Key, lc.rankTail.Score)
		fmt.Println("----")
	}
}

func (lc *LFUCache) delTail() {
	node := lc.rankTail
	if node == nil {
		return
	}
	delete(lc.cache, node.Key)
	delete(lc.keyRank, node.Key)

	if node.Prev != nil {
		node.Prev.Next = nil
	}
	lc.rankTail = node.Prev
}
