package main

type LockingTree struct {
	// 并不需要特别复杂的结构体..

	parent       []int
	lockNodeUser []int
	children     [][]int
}

func Constructor(parent []int) LockingTree {
	n := len(parent)
	lockNodeUser := make([]int, n)
	children := make([][]int, n)
	for i := 0; i < n; i++ {
		lockNodeUser[i] = -1
		p := parent[i]
		if p != -1 {
			children[p] = append(children[p], i)
		}
	}
	return LockingTree{
		parent:       parent,
		lockNodeUser: lockNodeUser,
		children:     children,
	}
}

func (lt *LockingTree) Lock(num int, user int) bool {
	if lt.lockNodeUser[num] == -1 {
		lt.lockNodeUser[num] = user
		return true
	}
	return false
}

func (lt *LockingTree) Unlock(num int, user int) bool {
	if lt.lockNodeUser[num] == user {
		lt.lockNodeUser[num] = -1
		return true
	}
	return false
}

func (lt *LockingTree) Upgrade(num int, user int) bool {
	res := lt.lockNodeUser[num] == -1 && !lt.hasLockedAncestor(num) && lt.checkAndUnlockDescendant(num)
	if res {
		lt.lockNodeUser[num] = user
	}
	return res
}

func (lt *LockingTree) hasLockedAncestor(num int) bool {
	num = lt.parent[num]
	for num != -1 {
		if lt.lockNodeUser[num] != -1 {
			return true
		}
		num = lt.parent[num]
	}
	return false
}

func (lt *LockingTree) checkAndUnlockDescendant(num int) bool {
	res := false
	// 只要有一个子节点(或者自己)被锁了, 那么就是可以解锁的
	if lt.lockNodeUser[num] != -1 {
		res = true
	}
	// 这点就很有意思: 为什么不需要处理回滚呢?
	// 因为如果没有锁, 那么本来就是-1,
	// 如果一旦有锁, 那么就应该是-1
	lt.lockNodeUser[num] = -1
	for _, child := range lt.children[num] {
		// 如果所有的子节点都没有锁, 那么就是false
		if lt.checkAndUnlockDescendant(child) {
			res = true
		}
	}
	return res
}
