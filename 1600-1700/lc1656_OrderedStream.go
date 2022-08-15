//go:build ignore

package main

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n),
		ptr:    1,
	}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	// fmt.Println(s.stream)
	s.stream[idKey-1] = value
	// fmt.Println(idKey, s.ptr)
	if idKey != s.ptr {
		return nil
	}
	end := s.ptr
	for end <= len(s.stream) {
		if s.stream[end-1] == "" {
			break
		}
		end++
	}
	var ret = s.stream[s.ptr-1 : end-1]
	s.ptr = end
	return ret
}
