package main

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(v string) {
	s.m[v] = exists
}

func (s *set) Remove(v string) {
	delete(s.m, v)
}

func (s *set) Contains(v string) bool {
	_, c := s.m[v]
	return c
}

func (s *set) Values() []string {
	var vs []string
	for v := range s.m {
		vs = append(vs, v)
	}
	return vs
}
