//go:build ignore

package main

type Node struct {
	Child  [26]*Node
	Finish bool
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
	node.Finish = true
}

func (n *Node) Search(word string, modify bool) bool {
	// 通过Modify锁定最多只能改一次字符
	if n == nil {
		return false
	}
	// 如果到了结尾, 必须要保证不相等(修改过字符)
	if len(word) == 0 {
		return modify && n.Finish
	}
	var c = int(word[0] - 'a')

	// 修改后边的和修改当前的
	if n.Child[c].Search(word[1:], modify) {
		return true
	}
	if !modify {
		// 尝试替换一下首字母试试
		for i, n := range n.Child {
			if i == c {
				continue
			}
			if n.Search(word[1:], true) {
				return true
			}
		}
	}
	return false
}

type MagicDictionary struct {
	root Node
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, str := range dictionary {
		this.root.Add(str)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.root.Search(searchWord, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
