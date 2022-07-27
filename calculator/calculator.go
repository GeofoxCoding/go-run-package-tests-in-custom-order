package calculator

var Pi = 3.141592654

func Sum(left, right int) int {
	sum := left + right
	return sum
}

func Divide(left, right int) (quotient int, remainder int) {
	quotient = left / right
	remainder = left % right

	return // Naked return, automatically returns all output parameters
}
