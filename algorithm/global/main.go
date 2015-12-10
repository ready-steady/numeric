// Package global provides an algorithm for globally adaptive hierarchical
// interpolation.
package global

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/internal"
)

var (
	infinity = math.Inf(1.0)
)

// Basis is a functional basis.
type Basis interface {
	// Compute evaluates the value of a basis function.
	Compute([]uint64, []float64) float64
}

// Grid is a sparse grid.
type Grid interface {
	// Compute returns the nodes corresponding to a set of indices.
	Compute([]uint64) []float64

	// Index returns the indices of a set of levels.
	Index([]uint8) []uint64
}

// Interpolator is an instance of the algorithm.
type Interpolator struct {
	config Config
	basis  Basis
	grid   Grid
}

// Progress contains information about the interpolation process.
type Progress struct {
	Level       uint8 // Reached level
	Active      uint  // The number of active indices
	Passive     uint  // The number of passive indices
	Evaluations uint  // The number of function evaluations
}

type cursor map[uint]bool

// New creates an interpolator.
func New(grid Grid, basis Basis, config *Config) *Interpolator {
	return &Interpolator{
		config: *config,
		basis:  basis,
		grid:   grid,
	}
}

// Compute constructs an interpolant for a function.
func (self *Interpolator) Compute(target Target) *Surrogate {
	config := &self.config

	ni, no := target.Dimensions()
	nw := config.Workers

	lindices := make([]uint8, 1*ni)
	indices := self.grid.Index(lindices)
	counts := []uint{uint(len(indices)) / ni}
	nodes := self.grid.Compute(indices)

	active := make(cursor)
	active[0] = true

	values := internal.Invoke(target.Compute, nodes, ni, no, nw)
	surrogate := newSurrogate(ni, no)
	surrogate.push(indices, values)

	terminator := newTerminator(no, config)
	terminator.push(values, values, counts)

	tracker := newTracker(ni, config)
	tracker.push(lindices, assess(target.Score, values, counts, no))

	progress := &Progress{Active: 1, Evaluations: counts[0]}

	for !terminator.done(active) {
		target.Monitor(progress)

		lindices = tracker.pull(active)
		nn := uint(len(lindices)) / ni

		progress.Active--
		progress.Passive++
		progress.Active += nn

		indices, counts = indices[:0], counts[:0]
		for i := uint(0); i < nn; i++ {
			newIndices := self.grid.Index(lindices[i*ni : (i+1)*ni])
			indices = append(indices, newIndices...)
			counts = append(counts, uint(len(newIndices))/ni)
		}

		level := maxUint8(lindices)
		if level > progress.Level {
			progress.Level = level
		}

		nn = uint(len(indices)) / ni
		progress.Evaluations += nn
		if progress.Evaluations > config.MaxEvaluations {
			break
		}

		nodes := self.grid.Compute(indices)
		values := internal.Invoke(target.Compute, nodes, ni, no, nw)
		surpluses := internal.Subtract(values, internal.Approximate(self.basis,
			surrogate.Indices, surrogate.Surpluses, nodes, ni, no, nw))

		surrogate.push(indices, surpluses)
		terminator.push(values, surpluses, counts)
		tracker.push(nil, assess(target.Score, surpluses, counts, no))
	}

	return surrogate
}

// Evaluate computes the values of an interpolant at a set of points.
func (self *Interpolator) Evaluate(surrogate *Surrogate, points []float64) []float64 {
	return internal.Approximate(self.basis, surrogate.Indices, surrogate.Surpluses, points,
		surrogate.Inputs, surrogate.Outputs, self.config.Workers)
}
