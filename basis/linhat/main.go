// Package linhat provides means for working with multi-dimensional bases
// formed by the linear hat function.
package linhat

import (
	"math"
)

// Self represents a particular instantiation of the basis.
type Self struct {
	dc uint16
}

// New creates an instance of the basis.
func New(dimensions uint16) *Self {
	return &Self{dimensions}
}

// Evaluate computes the value of a multi-dimensional basis function,
// identified by the given index, at the given point. Each element of the index
// is a pair (level, order) encoded into a single uint64.
func (self *Self) Evaluate(index []uint64, point []float64) float64 {
	var value, limit, delta float64 = 1, 0, 0
	var level uint8

	for i := uint16(0); i < self.dc; i++ {
		level = uint8(index[i] >> 32)
		if level == 0 {
			if math.Abs(point[i]-0.5) > 0.5 {
				return 0
			} else {
				continue
			}
		}

		limit = float64(uint32(2) << (level - 1))
		delta = math.Abs(point[i] - float64(uint32(index[i]))/limit)

		if delta >= 1/limit {
			return 0
		}

		value *= 1 - limit*delta
	}

	return value
}
