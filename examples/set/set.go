package set

// IntSet is a container of unique, unordered ints.
//
// It is a reference to a shared object so copies will reference the same
// underlying object. Clone() should be used to create a new set with the same
// elements.
type IntSet struct {
	m map[int]struct{}
}

// Len returns the number of elements in the set.
func (s IntSet) Len() int {
	return len(s.m)
}

// Insert inserts a value into the set. It returns true if a new element was
// inserted and false if the set already contained the value.
func (s *IntSet) Insert(value int) bool {
	if s.m == nil {
		s.m = make(map[int]struct{})
	}
	exists := s.Contains(value)
	s.m[value] = struct{}{}
	return !exists
}

// Contains checks whether the set contains the specified value.
func (s IntSet) Contains(value int) bool {
	_, ok := s.m[value]
	return ok
}

// Clone creates a new, independent copy of the set.
func (s IntSet) Clone() IntSet {
	if s.m == nil {
		return s
	}
	cloned := IntSet{
		m: make(map[int]struct{}, len(s.m)),
	}
	for v := range s.m {
		cloned.m[v] = struct{}{}
	}
	return cloned
}
