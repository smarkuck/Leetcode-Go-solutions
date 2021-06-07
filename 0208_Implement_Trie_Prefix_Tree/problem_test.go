package leetcode

import "testing"

type Trie struct {
	isEnd bool
	chars [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{isEnd: false, chars: [26]*Trie{}}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		char := c - 'a'
		if node.chars[char] == nil {
			node.chars[char] = &Trie{isEnd: false, chars: [26]*Trie{}}
		}
		node = node.chars[char]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(word string) *Trie {
	node := t
	for _, c := range word {
		char := c - 'a'
		if node.chars[char] != nil {
			node = node.chars[char]
		} else {
			return nil
		}
	}
	return node
}

func TestTrie(t *testing.T) {
	trie := Trie{isEnd: false, chars: [26]*Trie{}}
	trie.Insert("apple")

	output := trie.Search("apple")
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = trie.Search("app")
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = trie.StartsWith("app")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	trie.Insert("app")

	output = trie.Search("app")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
