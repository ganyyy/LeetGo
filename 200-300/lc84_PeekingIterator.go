package main

import "math"

type Iterator struct {
}

func (i *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
}
func (i *Iterator) next() int {
	// Returns the next element in the iteration.
}

type PeekingIterator struct {
	buf  int
	iter *Iterator
}

func Constructor44(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
		buf:  math.MaxInt32,
	}
}

func (pi *PeekingIterator) hasNext() bool {
	return pi.buf != math.MaxInt32 || pi.iter.hasNext()
}

func (pi *PeekingIterator) next() int {
	if pi.buf != math.MaxInt32 {
		var val = pi.buf
		pi.buf = math.MaxInt32
		return val
	}
	return pi.iter.next()
}

func (pi *PeekingIterator) peek() int {
	if pi.buf == math.MaxInt32 {
		pi.buf = pi.iter.next()
	}
	return pi.buf
}

type PeekingIterator2 struct {
	Iterator
}

func Constructor3(iter *Iterator) *PeekingIterator2 {
	return &PeekingIterator2{*iter}
}

func (it PeekingIterator2) peek() int {
	return it.next()
}
