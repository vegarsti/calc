package calc

type Operator func(left int, right int) int

var Plus = Operator(func(left int, right int) int { return left + right })
var Minus = Operator(func(left int, right int) int { return left - right })
var Times = Operator(func(left int, right int) int { return left * right })
var Divide = Operator(func(left int, right int) int { return left / right })
