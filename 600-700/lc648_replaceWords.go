//go:build ignore

package main

import "strings"

type Node struct {
	Child [26]*Node
	Str   string
}

func (n *Node) Add(word string) {
	var node = n
	for i := range word {
		var ch = word[i] - 'a'
		var child = node.Child[ch]
		if child == nil {
			child = &Node{}
			node.Child[ch] = child
		}
		node = child
	}
	node.Str = word
}

func (n *Node) Search(word string) string {
	var node = n
	for i := range word {
		var ch = word[i] - 'a'
		var child = node.Child[ch]
		if child == nil {
			return word
		}
		if child.Str != "" {
			return child.Str
		}
		node = child
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	// 构建字典树?

	var root = &Node{}
	for _, word := range dictionary {
		root.Add(word)
	}
	var words = strings.Split(sentence, " ")
	for i := range words {
		words[i] = root.Search(words[i])
	}

	return strings.Join(words, " ")
}
