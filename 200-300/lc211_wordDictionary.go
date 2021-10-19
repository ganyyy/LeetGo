package main

type Node struct {
	Next  [26]*Node
	IsEnd bool
}

func (n *Node) Search(w string) bool {
	if len(w) == 0 && n != nil {
		return n.IsEnd
	}
	if n == nil {
		return false
	}
	if w[0] == '.' {
		for _, next := range n.Next {
			if next != nil && next.Search(w[1:]) {
				return true
			}
		}
	} else {
		return n.Next[w[0]-'a'].Search(w[1:])
	}
	return false
}

type WordDictionary struct {
	// Trie?
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		Root: &Node{},
	}
}

func (d *WordDictionary) AddWord(word string) {
	var cur = d.Root
	for i := range word {
		var bs = word[i] - 'a'
		var next = cur.Next[bs]
		if next == nil {
			next = &Node{}
			cur.Next[bs] = next
		}
		cur = next
	}
	cur.IsEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	return d.Root.Search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
