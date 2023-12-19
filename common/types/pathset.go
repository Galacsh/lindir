package types

type PathSet map[string]struct{}

func (ps PathSet) String() string {
	var paths string

	for p := range ps {
		paths += p + "\n"
	}

	return paths
}

func (ps PathSet) Equals(other PathSet) bool {
	if len(ps) != len(other) {
		return false
	}

	for p := range ps {
		if !other.ContainsStr(p) {
			return false
		}
	}

	return true
}

func (ps PathSet) AddStr(s string) {
	ps[s] = struct{}{}
}

func (ps PathSet) Add(p Path) {
	ps.AddStr(p.String())
}

func (ps PathSet) RemoveStr(s string) {
	delete(ps, s)
}

func (ps PathSet) Remove(p Path) {
	ps.RemoveStr(p.String())
}

func (ps PathSet) ContainsStr(s string) bool {
	_, ok := ps[s]
	return ok
}

func (ps PathSet) Contains(p Path) bool {
	return ps.ContainsStr(p.String())
}

func (ps PathSet) Difference(other PathSet) PathSet {
	diff := PathSet{}

	for p := range ps {
		if !other.ContainsStr(p) {
			diff.AddStr(p)
		}
	}

	return diff
}

func (ps PathSet) Union(other PathSet) PathSet {
	union := PathSet{}

	for p := range ps {
		union.AddStr(p)
	}

	for p := range other {
		union.AddStr(p)
	}

	return union
}

func (ps PathSet) StrSlice() []string {
	var paths []string
	for p := range ps {
		paths = append(paths, p)
	}
	return paths
}

func (ps PathSet) Slice() []Path {
	var paths []Path
	for p := range ps {
		paths = append(paths, Path(p))
	}
	return paths
}

func (ps PathSet) Len() int {
	return len(ps)
}
