package hybrid

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/external"
)

// Target is a function to be interpolated.
type Target interface {
	// Dimensions returns the number of inputs and outputs.
	Dimensions() (uint, uint)

	// Done checks if the stopping criteria have been satisfied.
	Done(*external.Progress) bool

	// Compute evaluates the target function at a point.
	Compute([]float64, []float64)

	// Score assigns a score to an interpolation element.
	Score(*Element) []float64
}

// Element contains information about an interpolation element.
type Element struct {
	Lindex    []uint64  // Level index
	Indices   []uint64  // Nodal indices
	Volumes   []float64 // Basis-function volumes
	Values    []float64 // Target-function values
	Surpluses []float64 // Hierarchical surpluses
}

// BasicTarget is a basic target satisfying the Target interface.
type BasicTarget struct {
	DoneHandler    func(*external.Progress) bool
	ComputeHandler func([]float64, []float64) // != nil
	ScoreHandler   func(*Element) []float64

	ni uint
	no uint

	lmin uint
	lmax uint
}

// NewTarget creates a basic target.
func NewTarget(inputs, outputs uint, config *Config,
	compute func([]float64, []float64)) *BasicTarget {

	return &BasicTarget{
		ComputeHandler: compute,

		ni: inputs,
		no: outputs,

		lmin: config.MinLevel,
		lmax: config.MaxLevel,
	}
}

func (self *BasicTarget) Dimensions() (uint, uint) {
	return self.ni, self.no
}

func (self *BasicTarget) Done(progress *external.Progress) bool {
	if self.DoneHandler != nil {
		return self.DoneHandler(progress)
	} else {
		return false
	}
}

func (self *BasicTarget) Compute(node, value []float64) {
	self.ComputeHandler(node, value)
}

func (self *BasicTarget) Score(element *Element) []float64 {
	if self.ScoreHandler != nil {
		return self.ScoreHandler(element)
	}

	no, local := self.no, make([]float64, len(element.Volumes))
	for i, surplus := range element.Surpluses {
		j := uint(i) % no
		local[j] = math.Max(local[j], math.Abs(element.Volumes[j]*surplus))
	}
	return local
}
