package main

type node struct {
	next  [26]*node
	isEnd bool
}

func newNode() *node {
	return &node{}
}

type Trie struct {
	*node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		node: newNode(),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var cur = this.node
	for i := range word {
		var b = word[i] - 'a'
		var n = cur.next[b]
		if n == nil {
			n = newNode()
			cur.next[b] = n
		}
		cur = n
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var cur = this.node
	for i := range word {
		if cur = cur.next[word[i]-'a']; cur == nil {
			return false
		}
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var cur = this.node
	for i := range prefix {
		if cur = cur.next[prefix[i]-'a']; cur == nil {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
