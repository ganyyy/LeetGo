package main

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { panic("") }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { panic("") }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { panic("") }

type NestedIterator struct {
	cur int
	val []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	// 写个递归呗, 全部塞进来
	var val []int
	var dfs func(nested *NestedInteger)
	dfs = func(nested *NestedInteger) {
		if nested == nil {
			return
		}
		if nested.IsInteger() {
			val = append(val, nested.GetInteger())
			return
		}
		for _, n := range nested.GetList() {
			dfs(n)
		}
	}

	for _, n := range nestedList {
		dfs(n)
	}
	return &NestedIterator{val: val}
}

func (this *NestedIterator) Next() int {
	this.cur++
	return this.val[this.cur-1]
}

func (this *NestedIterator) HasNext() bool {
	return this.cur < len(this.val)
}
