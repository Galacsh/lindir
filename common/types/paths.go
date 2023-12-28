package types

type Paths []Path

func (p *Paths) Add(path string) {
	*p = append(*p, Path(path))
}

func (p Paths) Concat(other Paths) Paths {
	return append(p, other...)
}
