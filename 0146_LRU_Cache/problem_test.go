package leetcode

import "testing"

type Node struct {
	next, prev *Node
	key, value int
}

type LRUCache struct {
	capacity   int
	nodes      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{next: nil, prev: nil, key: 0, value: 0}
	tail := &Node{next: nil, prev: head, key: 0, value: 0}
	head.next = tail
	return LRUCache{capacity: capacity, nodes: map[int]*Node{}, head: head, tail: tail}
}

func (l *LRUCache) moveToHead(node *Node) {
	removeNode(node)
	l.addNode(node)
}

func (l *LRUCache) popLastNode() int {
	last := l.tail.prev
	removeNode(last)
	return last.key
}

func (l *LRUCache) addNode(node *Node) {
	next := l.head.next
	node.prev = l.head
	node.next = next
	next.prev = node
	l.head.next = node
}

func removeNode(node *Node) {
	next, prev := node.next, node.prev
	prev.next = next
	next.prev = prev
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.nodes[key]; ok {
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.nodes[key]; ok {
		node.value = value
		l.moveToHead(node)
		return
	}
	if len(l.nodes) == l.capacity {
		delete(l.nodes, l.popLastNode())
	}
	node := &Node{next: nil, prev: nil, key: key, value: value}
	l.nodes[key] = node
	l.addNode(node)
}

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	output := cache.Get(1)
	expected := 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	cache.Put(3, 3)
	output = cache.Get(2)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	cache.Put(4, 4)
	output = cache.Get(1)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = cache.Get(3)
	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = cache.Get(4)
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
