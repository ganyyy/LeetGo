package main

import "math/rand"

const maxLevel = 32
const pFactor = 0.25

type SkipListNode struct {
	val     int
	forward []*SkipListNode
}

type SkipList struct {
	head *SkipListNode
	// 简单而言, 最底层(0)存储了所有的值
	// 除此之外, 每一层存储了部分值
	// 按照概率学统计, 层次越高, 出现的几率越小
	level int
}

func Constructor() SkipList {
	// 初始化一个头结点
	return SkipList{&SkipListNode{-1, make([]*SkipListNode, maxLevel)}, 0}
}

func (*SkipList) randomLevel() int {
	lv := 1
	// 0.25的概率随机到下一层, 上限是maxLevel
	// 层级越高随机到的几率越低
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (s *SkipList) Search(target int) bool {
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 target 的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于 target
	return curr != nil && curr.val == target
}

func (s *SkipList) Add(num int) {
	// Add需要找前缀
	update := make([]*SkipListNode, maxLevel)
	for i := range update {
		update[i] = s.head
	}
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := s.randomLevel()
	s.level = max(s.level, lv)
	newNode := &SkipListNode{num, make([]*SkipListNode, lv)}
	for i, node := range update[:lv] {
		// 对第 i 层的状态进行更新，将当前元素的 forward 指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (s *SkipList) Erase(num int) bool {
	update := make([]*SkipListNode, maxLevel)
	curr := s.head
	for i := s.level - 1; i >= 0; i-- {
		// 找到第 i 层小于且最接近 num 的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果值不存在则返回 false
	if curr == nil || curr.val != num {
		return false
	}
	for i := 0; i < s.level && update[i].forward[i] == curr; i++ {
		// 对第 i 层的状态进行更新，将 forward 指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的 level
	for s.level > 1 && s.head.forward[s.level-1] == nil {
		s.level--
	}
	return true
}
