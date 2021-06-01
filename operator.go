package calc

import "math"

type Operator func(left int, right int) int

var Plus = Operator(func(left int, right int) int { return left + right })
var Minus = Operator(func(left int, right int) int { return left - right })
var Multiply = Operator(func(left int, right int) int { return left * right })
var Divide = Operator(func(left int, right int) int { return left / right })
var Exponentiate = Operator(func(left int, right int) int { return int(math.Pow(float64(left), float64(right))) })
