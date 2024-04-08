package main

type Set struct {
	setMap map[string]bool
}

func NewSet() *Set {
	return &Set{setMap: make(map[string]bool)}
}

func NewSetWithArray(a []string) *Set {
	s := NewSet()
	for _, v := range a {
		s.setMap[v] = true
	}
	return s
}

func (s *Set) IsExists(str string) bool {
	return s.setMap[str]
}
