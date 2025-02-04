package polyfit

import (
	"fmt"
	"math"

	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/mat"
)

func polyfit(xs, ys []float64, degree int) ([]float64, error) {
	if len(xs) != len(ys) {
		return nil, fmt.Errorf("polyfit: len(xs) != len(ys): %d != %d", len(xs), len(ys))
	}
	if len(xs) == 0 {
		return nil, fmt.Errorf("polyfit: len(xs) == len(ys) == 0")
	}

	// https://en.wikipedia.org/wiki/Least_squares#Linear_least_squares
	X := mat.NewDense(len(xs), degree+1, nil)
	for i, x := range xs {
		for j := range degree + 1 {
			X.Set(i, j, math.Pow(x, float64(j)))
		}
	}

	Xt := X.T()

	XtX := new(mat.Dense)
	XtX.Mul(Xt, X)

	Y := mat.NewVecDense(len(ys), ys)

	XtY := new(mat.VecDense)
	XtY.MulVec(Xt, Y)

	var coeff mat.VecDense
	// TODO: check error ?
	_ = coeff.SolveVec(XtX, XtY)

	var result []float64
	for i := range degree + 1 {
		result = append(result, coeff.AtVec(i))
	}

	slices.Reverse(result)
	return result, nil
}
