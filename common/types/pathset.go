package types

type PathSet map[string]struct{}

func (ps PathSet) String() string {
	var paths string

	for p := range ps {
		paths += p + "\n"
	}

	return paths
}

func (ps PathSet) Add(s string) {
	ps[s] = struct{}{}
}

func (ps PathSet) AddPath(p Path) {
	ps.Add(p.String())
}

func (ps PathSet) Remove(s string) {
	delete(ps, s)
}

func (ps PathSet) RemovePath(p Path) {
	ps.Remove(p.String())
}

func (ps PathSet) Contains(s string) bool {
	_, ok := ps[s]
	return ok
}

func (ps PathSet) ContainsPath(p Path) bool {
	return ps.Contains(p.String())
}

func (ps PathSet) Union(other PathSet) PathSet {
	union := PathSet{}

	for p := range ps {
		union.Add(p)
	}

	for p := range other {
		union.Add(p)
	}

	return union
}

func (ps PathSet) Difference(other PathSet) PathSet {
	diff := PathSet{}

	for p := range ps {
		if !other.Contains(p) {
			diff.Add(p)
		}
	}

	return diff
}

func (ps PathSet) Len() int {
	return len(ps)
}
