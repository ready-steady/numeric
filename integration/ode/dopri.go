package ode

import (
	"errors"
	"math"
)

// DormandPrince is an integrator based on the Dormand–Prince method.
//
// https://en.wikipedia.org/wiki/Dormand–Prince_method
type DormandPrince struct {
	config Config
}

// NewDormandPrince creates a new Dormand–Prince integrator.
func NewDormandPrince(config *Config) (*DormandPrince, error) {
	if err := config.verify(); err != nil {
		return nil, err
	}
	return &DormandPrince{config: *config}, nil
}

// Compute integrates the system of differential equations y' = f(x, y). The
// derivative function should evaluate f(x, y) given x and y in its first and
// second arguments, respectively, and store the result in its third argument.
// The initial value of y is given by initial. The first and last elements of
// the points array specify the interval of integration; hence, points should
// contain at least two elements. If points contain more than two elements, the
// solution is returned at exactly those points; otherwise, apart from the
// endpoints, the solution is also given at a number of intermediary points that
// the algorithm internally goes through. The solution is returned in the first
// output of the function while the corresponding points in the second one.
func (self *DormandPrince) Compute(derivative func(float64, []float64, []float64),
	points []float64, initial []float64) ([]float64, []float64, *Stats, error) {

	const (
		c2 = 1.0 / 5
		c3 = 3.0 / 10
		c4 = 4.0 / 5
		c5 = 8.0 / 9

		a21 = 1.0 / 5
		a31 = 3.0 / 40
		a32 = 9.0 / 40
		a41 = 44.0 / 45
		a42 = -56.0 / 15
		a43 = 32.0 / 9
		a51 = 19372.0 / 6561
		a52 = -25360.0 / 2187
		a53 = 64448.0 / 6561
		a54 = -212.0 / 729
		a61 = 9017.0 / 3168
		a62 = -355.0 / 33
		a63 = 46732.0 / 5247
		a64 = 49.0 / 176
		a65 = -5103.0 / 18656
		a71 = 35.0 / 384
		a73 = 500.0 / 1113
		a74 = 125.0 / 192
		a75 = -2187.0 / 6784
		a76 = 11.0 / 84

		e1 = 71.0 / 57600
		e3 = -71.0 / 16695
		e4 = 71.0 / 1920
		e5 = -17253.0 / 339200
		e6 = 22.0 / 525
		e7 = -1.0 / 40

		power = 1.0 / 5
	)

	np, nd, nc := len(points), len(initial), 0

	stats := &Stats{}

	z := make([]float64, nd)

	y := make([]float64, nd)
	ynew := make([]float64, nd)

	f := make([]float64, 7*nd)
	f1 := f[0*nd : 1*nd]
	f2 := f[1*nd : 2*nd]
	f3 := f[2*nd : 3*nd]
	f4 := f[3*nd : 4*nd]
	f5 := f[4*nd : 5*nd]
	f6 := f[5*nd : 6*nd]
	f7 := f[6*nd : 7*nd]

	x, xend := points[0], points[np-1]

	// Should the solution be returned at fixed points?
	fixed := np > 2

	// Prepare the first iteration.
	copy(y, initial)
	derivative(x, y, f1)
	stats.Evaluations++

	config := &self.config

	abserr, relerr := config.AbsError, config.RelError
	threshold := abserr / relerr

	// Compute the limits on the step size.
	hmax := config.MaxStep
	if hmax == 0 {
		hmax = 0.1 * (xend - x)
	}

	// Choose the initial step size.
	h := config.TryStep
	if h == 0 {
		h = points[1] - x
		if h > hmax {
			h = hmax
		}

		scale := 0.0
		for i := 0; i < nd; i++ {
			s := y[i]
			if s < 0 {
				s = -s
			}

			if s < threshold {
				s = threshold
			}

			s = f1[i] / s
			if s < 0 {
				s = -s
			}

			if s > scale {
				scale = s
			}
		}
		scale = scale / (0.8 * math.Pow(relerr, power))

		if h*scale > 1 {
			h = 1 / scale
		}
	}

	var values []float64
	if fixed {
		values = make([]float64, np*nd)
	} else {
		values = make([]float64, 0, 2*nd)
		points = make([]float64, 0, 2)
	}

	// Done with the first point.
	if fixed {
		copy(values, initial)
	} else {
		values = append(values, initial...)
		points = append(points, x)
	}
	nc += 1

	var xnew, ε float64

	done := false

	for {
		stats.Steps++

		hmin := 16 * epsilon(x)

		if h < hmin {
			h = hmin
		}
		if h > hmax {
			h = hmax
		}

		// Close to the end?
		if 1.1*h >= xend-x {
			h = xend - x
			done = true
		}

		rejected := false

		for {
			// Step 1
			for i := 0; i < nd; i++ {
				z[i] = y[i] + h*a21*f1[i]
			}

			// Step 2
			derivative(x+c2*h, z, f2)
			for i := 0; i < nd; i++ {
				z[i] = y[i] + h*(a31*f1[i]+a32*f2[i])
			}

			// Step 3
			derivative(x+c3*h, z, f3)
			for i := 0; i < nd; i++ {
				z[i] = y[i] + h*(a41*f1[i]+a42*f2[i]+a43*f3[i])
			}

			// Step 4
			derivative(x+c4*h, z, f4)
			for i := 0; i < nd; i++ {
				z[i] = y[i] + h*(a51*f1[i]+a52*f2[i]+a53*f3[i]+a54*f4[i])
			}

			// Step 5
			derivative(x+c5*h, z, f5)
			for i := 0; i < nd; i++ {
				z[i] = y[i] + h*(a61*f1[i]+a62*f2[i]+a63*f3[i]+a64*f4[i]+a65*f5[i])
			}

			// Step 6
			derivative(x+h, z, f6)
			for i := 0; i < nd; i++ {
				ynew[i] = y[i] + h*(a71*f1[i]+a73*f3[i]+a74*f4[i]+a75*f5[i]+a76*f6[i])
			}

			xnew = x + h
			h = xnew - x

			// Step 1
			derivative(xnew, ynew, f7)

			stats.Evaluations += 6

			// Compute the relative error.
			ε = 0
			for i := 0; i < nd; i++ {
				scale := y[i]
				if scale < 0 {
					scale = -scale
				}
				if ynew[i] > 0 {
					if ynew[i] > scale {
						scale = ynew[i]
					}
				} else {
					if -ynew[i] > scale {
						scale = -ynew[i]
					}
				}
				if scale < threshold {
					scale = threshold
				}

				e := e1*f1[i] + e3*f3[i] + e4*f4[i] + e5*f5[i] + e6*f6[i] + e7*f7[i]
				if e < 0 {
					e = -e
				}

				e = h * e / scale
				if e > ε {
					ε = e
				}
			}

			if ε <= relerr {
				break
			}

			stats.Rejections++

			if h <= hmin {
				return nil, nil, stats, errors.New("encountered a step-size underflow")
			}

			// Shrink the step size as the current one has been rejected.
			if rejected {
				h = 0.5 * h
			} else if scale := 0.8 * math.Pow(relerr/ε, power); scale > 0.1 {
				h = scale * h
			} else {
				h = 0.1 * h
			}

			if h < hmin {
				h = hmin
			}

			done = false
			rejected = true
		}

		if fixed {
			for nc < np {
				if xnew-points[nc] < 0 {
					break
				}

				if points[nc] == xnew {
					copy(values[nc*nd:(nc+1)*nd], ynew)
				} else {
					interpolate(x, y, f, h, points[nc], values[nc*nd:(nc+1)*nd])
				}

				nc++
			}
		} else {
			values = append(values, ynew...)
			points = append(points, xnew)
			nc++
		}

		if done {
			break
		}

		x = xnew
		copy(f1, f7)
		copy(y, ynew)

		if rejected {
			continue
		}

		// Compute a new step size.
		if scale := 1.25 * math.Pow(ε/relerr, power); scale > 0.2 {
			h = h / scale
		} else {
			h = 5 * h
		}
	}

	return values, points, stats, nil
}

func epsilon(x float64) float64 {
	if x < 0 {
		x = -x
	}
	return math.Nextafter(x, x+1) - x
}

func interpolate(x float64, y, f []float64, h, xnext float64, ynext []float64) {
	const (
		c11 = 1.0
		c12 = -183.0 / 64
		c13 = 37.0 / 12
		c14 = -145.0 / 128
		c32 = 1500.0 / 371
		c33 = -1000.0 / 159
		c34 = 1000.0 / 371
		c42 = -125.0 / 32
		c43 = 125.0 / 12
		c44 = -375.0 / 64
		c52 = 9477.0 / 3392
		c53 = -729.0 / 106
		c54 = 25515.0 / 6784
		c62 = -11.0 / 7
		c63 = 11.0 / 3
		c64 = -55.0 / 28
		c72 = 3.0 / 2
		c73 = -4.0
		c74 = 5.0 / 2
	)

	nd := len(y)

	s1 := (xnext - x) / h
	s2 := s1 * s1
	s3 := s1 * s2
	s4 := s1 * s3

	for i := 0; i < nd; i++ {
		f1 := f[0*nd+i]
		f3 := f[2*nd+i]
		f4 := f[3*nd+i]
		f5 := f[4*nd+i]
		f6 := f[5*nd+i]
		f7 := f[6*nd+i]

		ynext[i] = y[i] +
			h*s1*(c11*f1) +
			h*s2*(c12*f1+c32*f3+c42*f4+c52*f5+c62*f6+c72*f7) +
			h*s3*(c13*f1+c33*f3+c43*f4+c53*f5+c63*f6+c73*f7) +
			h*s4*(c14*f1+c34*f3+c44*f4+c54*f5+c64*f6+c74*f7)
	}
}