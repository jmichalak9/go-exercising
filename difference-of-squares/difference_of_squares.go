// Package diffsquares implements CollatzConjecture function
package diffsquares

import "math"

// SquareOfSum returns square of sum of first n natural numbers
func SquareOfSum(n int) int {
	x := float64(n)
	return int(math.Pow(x*(x+1)/2, float64(2)))
}

// SumOfSquares returns sum of squares of first n natural numbers
func SumOfSquares(n int) int {
	x := float64(n)
	return int(math.Pow(x, float64(3))/3 + math.Pow(x, float64(2))/2 + x/6)
}

// Difference returns difference between square of sum and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
