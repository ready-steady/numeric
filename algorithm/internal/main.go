// Package internal contains types and functions shared by the interpolation
// algorithms.
package internal

// Basis is a functional basis.
type Basis interface {
	// Compute evaluates the value of a basis function.
	Compute([]uint64, []float64) float64
}
