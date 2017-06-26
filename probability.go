package probability

import (
	"fmt"
	"math"
	"math/rand"
)

const SIGMA = 1.0
const MU = 0.0
const TOLERANCE = 0.00001

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

func InverseNormalCdf(p, mu, sigma, tolerance float64) float64 {
	if mu != MU || sigma != SIGMA {
		return mu + sigma*InverseNormalCdf(p, MU, SIGMA, tolerance)
	}
	lowZ := -10.0
	hiZ := 10.0
	var midZ, midP float64
	for hiZ-lowZ > tolerance {
		midZ = (lowZ + hiZ) / float64(2)
		midP = NormalCdf(midZ, MU, SIGMA)
		if midP < p {
			lowZ = midZ
		} else if midP > p {
			hiZ = midZ
		} else {
			break
		}
	}
	return midZ
}

func BernoulliTrial(p float64) (result int64) {
	if rand.Float64() < p {
		result = 1
	}
	return
}

func Binomial(p float64, n int) (result int64) {
	if !(0.0 <= p && p <= 1.0) {
		panic(fmt.Sprintf("Invalid probability p: %f", p))
	}
	if n <= 0 {
		panic(fmt.Sprintf("Invalid parameter n: %d", n))
	}
	for i := 1; i <= n; i++ {
		result += BernoulliTrial(p)
	}
	return
}
