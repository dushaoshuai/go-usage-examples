package mathx

import (
	"math"
)

// EvalPolynomial calculates the value of a polynomial at a given x.
// coefficients are in descending order, i.e., coefficients[i] corresponds to the coefficient of the (n-i-1)-th power of x.
func EvalPolynomial(coefficients []float64, x float64) float64 {
	var result float64
	n := len(coefficients)
	for i, coeff := range coefficients {
		exponent := float64(n - i - 1)
		result += coeff * math.Pow(x, exponent)
	}
	return result
}
