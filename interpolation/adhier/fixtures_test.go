package adhier

import (
	"math"
)

type fixture struct {
	surrogate *Surrogate
	levels    []uint32
	orders    []uint32
	points    []float64
	values    []float64
}

func init() {
	fixtureStep.prepare()
	fixtureHat.prepare()
	fixtureCube.prepare()
	fixtureBox.prepare()
}

func (f *fixture) prepare() {
	f.surrogate.Indices = make([]uint64, len(f.levels))
	for i := range f.levels {
		f.surrogate.Indices[i] = uint64(f.levels[i]) | uint64(f.orders[i])<<32
	}
}

func step(x, y []float64, _ []uint64) {
	if x[0] <= 0.5 {
		y[0] = 1
	} else {
		y[0] = 0
	}
}

var fixtureStep = fixture{
	surrogate: &Surrogate{
		Inputs:  1,
		Outputs: 1,

		Level: 4,
		Nodes: 8,

		Surpluses: []float64{1, 0, -1, -0.5, -0.5, 0, -0.5, 0},
	},
	levels: []uint32{0, 1, 1, 2, 3, 3, 4, 4},
	orders: []uint32{0, 0, 2, 3, 5, 7, 9, 11},
	points: []float64{0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1},
	values: []float64{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
}

func hat(x, y []float64, _ []uint64) {
	z := 5*x[0] - 1

	switch {
	case 0 <= z && z < 1:
		y[0] = 0.5 * z * z
	case 1 <= z && z < 2:
		y[0] = 0.5 * (-2*z*z + 6*z - 3)
	case 2 <= z && z < 3:
		y[0] = 0.5 * (3 - z) * (3 - z)
	default:
		y[0] = 0
	}
}

var fixtureHat = fixture{
	surrogate: &Surrogate{
		Inputs:  1,
		Outputs: 1,

		Level: 9,
		Nodes: 305,

		Surpluses: []float64{
			+7.50000000000000e-01, -7.50000000000000e-01, -7.50000000000000e-01,
			-3.43750000000000e-01, -3.43750000000000e-01, -1.56250000000000e-02,
			-7.81250000000000e-03, -7.81250000000000e-03, -1.56250000000000e-02,
			+0.00000000000000e+00, -1.56250000000000e-02, -4.88281250000000e-02,
			+8.59375000000000e-02, +8.59375000000000e-02, -4.88281250000000e-02,
			-1.56250000000000e-02, +0.00000000000000e+00, +0.00000000000000e+00,
			-1.12304687500000e-02, -1.22070312500000e-02, -1.22070312500000e-02,
			+1.26953125000000e-02, +2.44140625000000e-02, +2.44140625000000e-02,
			+1.26953125000000e-02, -1.22070312500000e-02, -1.22070312500000e-02,
			-1.12304687500000e-02, +0.00000000000000e+00, -2.07519531250000e-03,
			-3.05175781250000e-03, -3.05175781250000e-03, -3.05175781250000e-03,
			-3.05175781250000e-03, -3.05175781250000e-03, -2.31933593750000e-03,
			+6.10351562500000e-03, +6.10351562500000e-03, +6.10351562500000e-03,
			+6.10351562500000e-03, +6.10351562500000e-03, +6.10351562500000e-03,
			-2.31933593750000e-03, -3.05175781250000e-03, -3.05175781250000e-03,
			-3.05175781250000e-03, -3.05175781250000e-03, -3.05175781250000e-03,
			-2.07519531250000e-03, -6.10351562500000e-05, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -3.05175781250000e-05,
			+1.52587890625000e-03, +1.52587890625000e-03, +1.52587890625000e-03,
			+1.52587890625000e-03, +1.52587890625000e-03, +1.52587890625000e-03,
			+1.52587890625000e-03, +1.52587890625000e-03, +1.52587890625000e-03,
			+1.52587890625000e-03, +1.52587890625000e-03, +1.52587890625000e-03,
			-3.05175781250000e-05, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -7.62939453125000e-04, -7.62939453125000e-04,
			-7.62939453125000e-04, -6.10351562500000e-05, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, +3.81469726562500e-04,
			+3.81469726562500e-04, +3.81469726562500e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -1.90734863281250e-04,
			-1.90734863281250e-04, -1.90734863281250e-04, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, +9.53674316406250e-05,
			+9.53674316406250e-05, +9.53674316406250e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05, -4.76837158203125e-05,
			-4.76837158203125e-05, -4.76837158203125e-05,
		},
	},

	levels: []uint32{
		0, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
		6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7,
		7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
		8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
		8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
		8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
		8, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9,
		9, 9, 9, 9, 9,
	},

	orders: []uint32{
		0, 0, 2, 1, 3, 1, 3, 5, 7, 1,
		3, 5, 7, 9, 11, 13, 15, 5, 7, 9,
		11, 13, 15, 17, 19, 21, 23, 25, 27, 13,
		15, 17, 19, 21, 23, 25, 27, 29, 31, 33,
		35, 37, 39, 41, 43, 45, 47, 49, 51, 25,
		27, 29, 31, 33, 35, 37, 39, 41, 43, 45,
		47, 49, 51, 53, 55, 57, 59, 61, 63, 65,
		67, 69, 71, 73, 75, 77, 79, 81, 83, 85,
		87, 89, 91, 93, 95, 97, 99, 101, 103, 53,
		55, 57, 59, 61, 63, 65, 67, 69, 71, 73,
		75, 77, 79, 81, 83, 85, 87, 89, 91, 93,
		95, 97, 99, 105, 107, 109, 111, 113, 115, 117,
		119, 121, 123, 125, 127, 129, 131, 133, 135, 137,
		139, 141, 143, 145, 147, 149, 151, 157, 159, 161,
		163, 165, 167, 169, 171, 173, 175, 177, 179, 181,
		183, 185, 187, 189, 191, 193, 195, 197, 199, 201,
		203, 105, 107, 109, 111, 113, 115, 117, 119, 121,
		123, 125, 127, 129, 131, 133, 135, 137, 139, 141,
		143, 145, 147, 149, 151, 153, 155, 157, 159, 161,
		163, 165, 167, 169, 171, 173, 175, 177, 179, 181,
		183, 185, 187, 189, 191, 193, 195, 197, 199, 209,
		211, 213, 215, 217, 219, 221, 223, 225, 227, 229,
		231, 233, 235, 237, 239, 241, 243, 245, 247, 249,
		251, 253, 255, 257, 259, 261, 263, 265, 267, 269,
		271, 273, 275, 277, 279, 281, 283, 285, 287, 289,
		291, 293, 295, 297, 299, 301, 303, 313, 315, 317,
		319, 321, 323, 325, 327, 329, 331, 333, 335, 337,
		339, 341, 343, 345, 347, 349, 351, 353, 355, 357,
		359, 361, 363, 365, 367, 369, 371, 373, 375, 377,
		379, 381, 383, 385, 387, 389, 391, 393, 395, 397,
		399, 401, 403, 405, 407,
	},

	points: []float64{
		0.00, 0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09,
		0.10, 0.11, 0.12, 0.13, 0.14, 0.15, 0.16, 0.17, 0.18, 0.19,
		0.20, 0.21, 0.22, 0.23, 0.24, 0.25, 0.26, 0.27, 0.28, 0.29,
		0.30, 0.31, 0.32, 0.33, 0.34, 0.35, 0.36, 0.37, 0.38, 0.39,
		0.40, 0.41, 0.42, 0.43, 0.44, 0.45, 0.46, 0.47, 0.48, 0.49,
		0.50, 0.51, 0.52, 0.53, 0.54, 0.55, 0.56, 0.57, 0.58, 0.59,
		0.60, 0.61, 0.62, 0.63, 0.64, 0.65, 0.66, 0.67, 0.68, 0.69,
		0.70, 0.71, 0.72, 0.73, 0.74, 0.75, 0.76, 0.77, 0.78, 0.79,
		0.80, 0.81, 0.82, 0.83, 0.84, 0.85, 0.86, 0.87, 0.88, 0.89,
		0.90, 0.91, 0.92, 0.93, 0.94, 0.95, 0.96, 0.97, 0.98, 0.99,
		1.00,
	},

	values: []float64{
		+0.0000000000000000e+00, +1.7347234759768071e-18, +3.9898639947466563e-17,
		+4.5102810375396984e-17, -6.5919492087118670e-17, -2.2551405187698492e-17,
		-2.0816681711721685e-17, -2.0816681711721685e-17, -1.9081958235744878e-17,
		+1.0408340855860843e-17, -1.0061396160665481e-16, +3.9898639947466563e-17,
		-4.1633363423443370e-17, +4.1633363423443370e-17, -4.1633363423443370e-17,
		+1.0061396160665481e-16, -1.0408340855860843e-17, +4.5102810375396984e-17,
		+2.0816681711721685e-17, +4.8572257327350599e-17, +7.3242187500020990e-05,
		+1.2619018554686885e-03, +5.0109863281249702e-03, +1.1258697509765575e-02,
		+2.0005035400390642e-02, +3.1250000000000000e-02, +4.5005035400390664e-02,
		+6.1258697509765647e-02, +8.0010986328125019e-02, +1.0126190185546868e-01,
		+1.2501144409179682e-01, +1.5125961303710936e-01, +1.8000640869140624e-01,
		+2.1125183105468748e-01, +2.4500350952148442e-01, +2.8125762939453136e-01,
		+3.2001037597656251e-01, +3.6126174926757804e-01, +4.0501174926757821e-01,
		+4.5126037597656243e-01, +4.9982910156250021e-01, +5.4749298095703125e-01,
		+5.8999633789062500e-01, +6.2748718261718761e-01, +6.5998077392578125e-01,
		+6.8747711181640625e-01, +7.0997619628906261e-01, +7.2747802734375011e-01,
		+7.3998260498046875e-01, +7.4748992919921875e-01, +7.5000000000000000e-01,
		+7.4748992919921875e-01, +7.3998260498046875e-01, +7.2747802734375000e-01,
		+7.0997619628906239e-01, +6.8747711181640614e-01, +6.5998077392578125e-01,
		+6.2748718261718728e-01, +5.8999633789062478e-01, +5.4749298095703125e-01,
		+4.9982910156250021e-01, +4.5126037597656243e-01, +4.0501174926757821e-01,
		+3.6126174926757804e-01, +3.2001037597656251e-01, +2.8125762939453158e-01,
		+2.4500350952148470e-01, +2.1125183105468776e-01, +1.8000640869140647e-01,
		+1.5125961303710952e-01, +1.2501144409179704e-01, +1.0126190185546882e-01,
		+8.0010986328125019e-02, +6.1258697509765647e-02, +4.5005035400390664e-02,
		+3.1250000000000000e-02, +2.0005035400390642e-02, +1.1258697509765613e-02,
		+5.0109863281249589e-03, +1.2619018554687412e-03, +7.3242187500027062e-05,
		+0.0000000000000000e+00, +0.0000000000000000e+00, -1.3877787807814457e-17,
		-4.1633363423443370e-17, +4.1633363423443370e-17, -4.1633363423443370e-17,
		+4.1633363423443370e-17, -4.1633363423443370e-17, +4.1633363423443370e-17,
		-4.1633363423443370e-17, +5.5511151231257827e-17, -5.5511151231257827e-17,
		+0.0000000000000000e+00, +0.0000000000000000e+00, +0.0000000000000000e+00,
		+0.0000000000000000e+00, +0.0000000000000000e+00, +0.0000000000000000e+00,
		+0.0000000000000000e+00, +0.0000000000000000e+00,
	},
}

func cube(x, y []float64, _ []uint64) {
	if math.Abs(2*x[0]-1) < 0.45 && math.Abs(2*x[1]-1) < 0.45 {
		y[0] = 1
	} else {
		y[0] = 0
	}
}

var fixtureCube = fixture{
	surrogate: &Surrogate{
		Inputs:  2,
		Outputs: 1,

		Level: 3,
		Nodes: 29,

		Surpluses: []float64{
			1, -1, -1, -1, -1, -0.5, 1, 1, -0.5, 1, 1, -0.5, -0.5, 0, 0.5,
			0.5, 0.5, 0.5, 0.5, 0.5, 0, 0.5, 0.5, 0.5, 0.5, 0, 0.5, 0.5, 0,
		},
	},

	levels: []uint32{
		0, 0,
		1, 0,
		1, 0,
		0, 1,
		0, 1,
		2, 0,
		1, 1,
		1, 1,
		2, 0,
		1, 1,
		1, 1,
		0, 2,
		0, 2,
		3, 0,
		3, 0,
		2, 1,
		2, 1,
		1, 2,
		1, 2,
		3, 0,
		3, 0,
		2, 1,
		2, 1,
		1, 2,
		1, 2,
		0, 3,
		0, 3,
		0, 3,
		0, 3,
	},

	orders: []uint32{
		0, 0,
		0, 0,
		2, 0,
		0, 0,
		0, 2,
		1, 0,
		0, 0,
		0, 2,
		3, 0,
		2, 0,
		2, 2,
		0, 1,
		0, 3,
		1, 0,
		3, 0,
		1, 0,
		1, 2,
		0, 1,
		0, 3,
		5, 0,
		7, 0,
		3, 0,
		3, 2,
		2, 1,
		2, 3,
		0, 1,
		0, 3,
		0, 5,
		0, 7,
	},
}

func box(x, y []float64, _ []uint64) {
	if x[0]+x[1] > 0.5 {
		y[0] = 1
	} else {
		y[0] = 0
	}

	if x[0]-x[1] > 0.5 {
		y[1] = 1
	} else {
		y[1] = 0
	}

	if x[1]-x[0] > 0.5 {
		y[2] = 1
	} else {
		y[2] = 0
	}
}

var fixtureBox = fixture{
	surrogate: &Surrogate{
		Inputs:  2,
		Outputs: 3,

		Level: 3,
		Nodes: 20,

		Surpluses: []float64{
			+1.0, 0.0, 0.0,
			-1.0, 0.0, 0.0,
			+0.0, 0.0, 0.0,
			-1.0, 0.0, 0.0,
			+0.0, 0.0, 0.0,
			+0.5, 0.0, 0.0,
			+1.0, 0.0, 0.0,
			+1.0, 0.0, 1.0,
			+1.0, 1.0, 0.0,
			+0.5, 0.0, 0.0,
			+0.5, 0.0, 0.0,
			+0.0, 0.0, 0.0,
			-0.5, 0.0, 0.0,
			-0.5, 0.0, 0.5,
			-0.5, 0.0, 0.0,
			+0.5, 0.0, 0.5,
			+0.5, 0.5, 0.0,
			-0.5, 0.5, 0.0,
			+0.5, 0.0, 0.0,
			+0.0, 0.0, 0.0,
		},
	},

	levels: []uint32{
		0, 0,
		1, 0,
		1, 0,
		0, 1,
		0, 1,
		2, 0,
		1, 1,
		1, 1,
		1, 1,
		0, 2,
		3, 0,
		3, 0,
		2, 1,
		2, 1,
		1, 2,
		1, 2,
		2, 1,
		1, 2,
		0, 3,
		0, 3,
	},

	orders: []uint32{
		0, 0,
		0, 0,
		2, 0,
		0, 0,
		0, 2,
		1, 0,
		0, 0,
		0, 2,
		2, 0,
		0, 1,
		1, 0,
		3, 0,
		1, 0,
		1, 2,
		0, 1,
		0, 3,
		3, 0,
		2, 1,
		0, 1,
		0, 3,
	},

	points: []float64{
		0.0, 0.0,
		0.0, 0.1,
		0.0, 0.2,
		0.0, 0.3,
		0.0, 0.4,
		0.0, 0.5,
		0.0, 0.6,
		0.0, 0.7,
		0.0, 0.8,
		0.0, 0.9,
		0.0, 1.0,
		0.1, 0.0,
		0.1, 0.1,
		0.1, 0.2,
		0.1, 0.3,
		0.1, 0.4,
		0.1, 0.5,
		0.1, 0.6,
		0.1, 0.7,
		0.1, 0.8,
		0.1, 0.9,
		0.1, 1.0,
		0.2, 0.0,
		0.2, 0.1,
		0.2, 0.2,
		0.2, 0.3,
		0.2, 0.4,
		0.2, 0.5,
		0.2, 0.6,
		0.2, 0.7,
		0.2, 0.8,
		0.2, 0.9,
		0.2, 1.0,
		0.3, 0.0,
		0.3, 0.1,
		0.3, 0.2,
		0.3, 0.3,
		0.3, 0.4,
		0.3, 0.5,
		0.3, 0.6,
		0.3, 0.7,
		0.3, 0.8,
		0.3, 0.9,
		0.3, 1.0,
		0.4, 0.0,
		0.4, 0.1,
		0.4, 0.2,
		0.4, 0.3,
		0.4, 0.4,
		0.4, 0.5,
		0.4, 0.6,
		0.4, 0.7,
		0.4, 0.8,
		0.4, 0.9,
		0.4, 1.0,
		0.5, 0.0,
		0.5, 0.1,
		0.5, 0.2,
		0.5, 0.3,
		0.5, 0.4,
		0.5, 0.5,
		0.5, 0.6,
		0.5, 0.7,
		0.5, 0.8,
		0.5, 0.9,
		0.5, 1.0,
		0.6, 0.0,
		0.6, 0.1,
		0.6, 0.2,
		0.6, 0.3,
		0.6, 0.4,
		0.6, 0.5,
		0.6, 0.6,
		0.6, 0.7,
		0.6, 0.8,
		0.6, 0.9,
		0.6, 1.0,
		0.7, 0.0,
		0.7, 0.1,
		0.7, 0.2,
		0.7, 0.3,
		0.7, 0.4,
		0.7, 0.5,
		0.7, 0.6,
		0.7, 0.7,
		0.7, 0.8,
		0.7, 0.9,
		0.7, 1.0,
		0.8, 0.0,
		0.8, 0.1,
		0.8, 0.2,
		0.8, 0.3,
		0.8, 0.4,
		0.8, 0.5,
		0.8, 0.6,
		0.8, 0.7,
		0.8, 0.8,
		0.8, 0.9,
		0.8, 1.0,
		0.9, 0.0,
		0.9, 0.1,
		0.9, 0.2,
		0.9, 0.3,
		0.9, 0.4,
		0.9, 0.5,
		0.9, 0.6,
		0.9, 0.7,
		0.9, 0.8,
		0.9, 0.9,
		0.9, 1.0,
		1.0, 0.0,
		1.0, 0.1,
		1.0, 0.2,
		1.0, 0.3,
		1.0, 0.4,
		1.0, 0.5,
		1.0, 0.6,
		1.0, 0.7,
		1.0, 0.8,
		1.0, 0.9,
		1.0, 1.0,
	},

	values: []float64{
		0.00, 0.00, 0.00,
		0.40, 0.00, 0.00,
		0.20, 0.00, 0.00,
		0.00, 0.00, 0.00,
		0.00, 0.00, 0.00,
		0.00, 0.00, 0.00,
		0.40, 0.00, 0.40,
		0.80, 0.00, 0.80,
		1.00, 0.00, 1.00,
		1.00, 0.00, 1.00,
		1.00, 0.00, 1.00,
		0.40, 0.00, 0.00,
		0.92, 0.00, 0.00,
		0.84, 0.00, 0.00,
		0.72, 0.00, 0.00,
		0.76, 0.00, 0.00,
		0.80, 0.00, 0.00,
		1.08, 0.00, 0.36,
		1.36, 0.00, 0.72,
		1.48, 0.00, 0.92,
		1.44, 0.00, 0.96,
		1.40, 0.00, 1.00,
		0.20, 0.00, 0.00,
		0.84, 0.00, 0.00,
		0.88, 0.00, 0.00,
		0.84, 0.00, 0.00,
		0.92, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.16, 0.00, 0.32,
		1.32, 0.00, 0.64,
		1.36, 0.00, 0.84,
		1.28, 0.00, 0.92,
		1.20, 0.00, 1.00,
		0.00, 0.00, 0.00,
		0.72, 0.00, 0.00,
		0.84, 0.00, 0.00,
		0.84, 0.00, 0.00,
		0.92, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.08, 0.00, 0.24,
		1.16, 0.00, 0.48,
		1.16, 0.00, 0.64,
		1.08, 0.00, 0.72,
		1.00, 0.00, 0.80,
		0.00, 0.00, 0.00,
		0.76, 0.00, 0.00,
		0.92, 0.00, 0.00,
		0.92, 0.00, 0.00,
		0.96, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.04, 0.00, 0.12,
		1.08, 0.00, 0.24,
		1.08, 0.00, 0.32,
		1.04, 0.00, 0.36,
		1.00, 0.00, 0.40,
		0.00, 0.00, 0.00,
		0.80, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		0.40, 0.40, 0.00,
		1.08, 0.36, 0.00,
		1.16, 0.32, 0.00,
		1.08, 0.24, 0.00,
		1.04, 0.12, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		0.80, 0.80, 0.00,
		1.36, 0.72, 0.00,
		1.32, 0.64, 0.00,
		1.16, 0.48, 0.00,
		1.08, 0.24, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 1.00, 0.00,
		1.48, 0.92, 0.00,
		1.36, 0.84, 0.00,
		1.16, 0.64, 0.00,
		1.08, 0.32, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 1.00, 0.00,
		1.44, 0.96, 0.00,
		1.28, 0.92, 0.00,
		1.08, 0.72, 0.00,
		1.04, 0.36, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 1.00, 0.00,
		1.40, 1.00, 0.00,
		1.20, 1.00, 0.00,
		1.00, 0.80, 0.00,
		1.00, 0.40, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
		1.00, 0.00, 0.00,
	},
}

func many(ni, no int) func([]float64, []float64, []uint64) {
	return func(x, y []float64, _ []uint64) {
		sum, value := 0.0, 0.0

		for i := 0; i < ni; i++ {
			sum += x[i]
		}

		if sum > float64(ni)/4 {
			value = 1
		}

		for i := 0; i < no; i++ {
			y[i] = value
		}
	}
}
