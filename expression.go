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

type BinaryExpression struct {
	left     Expression
	right    Expression
	operator Operator
}

func NewBinaryExpression(left Expression, right Expression, operator Operator) BinaryExpression {
	return BinaryExpression{
		left:     left,
		right:    right,
		operator: operator,
	}
}

func (b BinaryExpression) Value() int {
	return b.operator(b.left.Value(), b.right.Value())
}

func (b BinaryExpression) Kind() Kind {
	return BinaryKind
}
