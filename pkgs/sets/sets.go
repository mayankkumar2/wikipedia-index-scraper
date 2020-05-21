package sets


var exists = struct{}{}

type Set struct {
	m map[string]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *Set) Add(value string) {
	s.m[value] = exists
}

func (s *Set) Remove(value string) {
	delete(s.m, value)
}

func (s *Set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}
func (s *Set) Slice() []string {
	c := make([]string,len(s.m))
	cnt := 0
	for i, _ := range s.m {
		c[cnt] = i
		cnt++
	}
	return c
}