package probability

import "math"

const SIGMA = 1.0
const MU = 0.0

func UniformPdf(x float64) (result int) {
	result = 0
	if x >= 0 && x < 1 {
		result = 1
	}
	return
}

func UniformCdf(x float64) (result float64) {
	result = 1
	if x < 0 {
		result = 0
	} else if x < 1 {
		result = x
	}
	return
}

func NormalPdf(x, mu, sigma float64) float64 {
	sqrtTwoPi := math.Sqrt(2 * math.Pi)
	powXmu := math.Pow((x - mu), 2)
	exp := math.Exp(-(powXmu / (2 * math.Pow(sigma, 2))))
	return exp / (sqrtTwoPi * sigma)
}

func NormalCdf(x, mu, sigma float64) float64 {
	erf := math.Erf((x - mu) / math.Sqrt2 / sigma)
	return (1 + erf) / float64(2)
}
