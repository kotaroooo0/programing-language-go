package ex07

type IntSet struct {
	words []uint64
}

func (s *IntSet) Elems() []int {
	r := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, 64*i+j)
			}
		}
	}
	return r
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, sword := range s.words {
		sword &= t.words[i]
	}
	s.words = s.words[:len(t.words)]
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	w := make([]uint64, len(s.words))
	copy(w, s.words)
	return &IntSet{w}
}
