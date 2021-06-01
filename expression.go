package calc

type Expression interface {
	Value() int
	Kind() Kind
}

type Atom struct {
	value int
}

func NewAtom(value int) Atom {
	return Atom{value: value}
}

func (a Atom) Value() int {
	return a.value
}

func (a Atom) Kind() Kind {
	return AtomKind
}
