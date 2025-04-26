package operation

import . "golang.org/x/exp/constraints"

// Number Define a custom interface that includes both Integer and Float
type Number interface {
	Integer | Float
}

// Add Bringing two or more numbers (or things) together to make a new total
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
// Generic Add function
func Add[T Number](a T, b T) T {
	return a + b
}
