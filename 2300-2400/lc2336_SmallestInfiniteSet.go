package main

type SmallestInfiniteSet struct {
	min  int
	data [126]uint8 // 对应bit为1表示已pop
}

func Constructor() SmallestInfiniteSet { return SmallestInfiniteSet{} }

func split(n int) (idx, offset uint8) {
	return uint8(n / 8), uint8(1 << (n & 7))
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	idx, offset := split(int(s.min))
	if s.data[idx]&offset == 0 {
		s.data[idx] |= offset
		s.min++
		return s.min
	}
	for i := s.min + 1; i < 1000; i++ {
		idx, offset := split(i)
		if s.data[idx]&offset == 0 {
			s.data[idx] |= offset
			s.min = i + 1
			return s.min
		}
	}
	return -1
}

func (s *SmallestInfiniteSet) AddBack(n int) {
	n--
	idx, offset := split(n)
	s.data[idx] &^= offset
	if n < s.min {
		s.min = n
	}
}
