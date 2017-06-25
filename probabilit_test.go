package probability

import (
	"testing"
)

func TestUniformPdf(t *testing.T) {
	cases := []struct {
		x    float64
		want int
	}{
		{0.65, 1},
		{1, 0},
		{-1, 0},
		{-0.4, 0},
	}

	for _, c := range cases {
		got := UniformPdf(c.x)
		if got != c.want {
			t.Errorf("UniformPdf(%f) want: %d but got: %d", c.x, c.want, got)
		}
	}
}

func TestUniformCdf(t *testing.T) {
	cases := []struct {
		x    float64
		want float64
	}{
		{1.0, 1.0},
		{0.3, 0.3},
		{-0.2, 0.0},
	}

	for _, c := range cases {
		got := UniformCdf(c.x)
		if got != c.want {
			t.Errorf("UniformCdf(%f) want: %f but got: %f", c.x, c.want, got)
		}
	}
}

func TestNormalPdf(t *testing.T) {
	cases := []struct {
		x, mu, sigma float64
		want         float64
	}{
		{1.0, MU, SIGMA, 0.24197072451914337},
		{2.0, MU, SIGMA, 0.05399096651318806},
		{12.0, MU, SIGMA, 2.1463837356630605e-32},
		{0.0, MU, SIGMA, 0.3989422804014327},
		{-10.0, MU, SIGMA, 7.69459862670642e-23},
		{-10.0, 0.0, SIGMA, 7.69459862670642e-23},
		{-10.0, 2, 1, 2.1463837356630605e-32},
		{-10.0, 2, 2, 3.037941424911643e-09},
		{-1.0, 2, 1, 0.0044318484119380075},
		{10.0, 2, 3, 0.0037986620079324806},
	}
	for _, c := range cases {
		got := NormalPdf(c.x, c.mu, c.sigma)
		if got != c.want {
			t.Errorf("NormalPdf(%v, %v, %v) want: %v but got: %v",
				c.x, c.sigma, c.mu, c.want, got)
		}
	}
}

func TestNormalCdf(t *testing.T) {
	cases := []struct {
		x, mu, sigma float64
		want         float64
	}{
		{10.0, MU, SIGMA, 1},
		{3.0, MU, SIGMA, 0.9986501019683699},
		{1.0, MU, SIGMA, 0.8413447460685429},
		{1.0, 1, 2, 0.5},
		{1.0, 2, 2, 0.3085375387259869},
		{1.0, -2, SIGMA, 0.9986501019683699},
		{3.0, 2, -2, 0.3085375387259869},
		{-5.0, -3, -2, 0.8413447460685429},
	}
	for _, c := range cases {
		got := NormalCdf(c.x, c.mu, c.sigma)
		if got != c.want {
			t.Errorf("NormalCdf(%v, %v, %v) want: %v but got: %v",
				c.x, c.mu, c.sigma, c.want, got)
		}
	}
}

func TestInverseNormalCdf(t *testing.T) {
	cases := []struct {
		p                    float64
		mu, sigma, tolerance float64
		want                 float64
	}{
		{1.0, MU, SIGMA, TOLERANCE, 8.75},
		{0.2, MU, SIGMA, TOLERANCE, -0.8416271209716797},
		{0.5, MU, SIGMA, TOLERANCE, 0.0},
		{1.0, 1, SIGMA, TOLERANCE, 9.75},
		{1.0, MU, 2, TOLERANCE, 17.5},
		{1.0, MU, SIGMA, 2 * TOLERANCE, 8.75},
		{1.0, MU, SIGMA, 0.1, 8.75},
		{1.0, 2, 2, 0.0, 19.5},
		{1.0, 2, 2, 2.0, 19.5},
		{0.0, 0, 0, 0.0, 0.0},
	}

	for _, c := range cases {
		got := InverseNormalCdf(c.p, c.mu, c.sigma, c.tolerance)
		if got != c.want {
			t.Errorf("InverseNormalCdf(%v,%v,%v,%v) want: %v but got: %v",
				c.p, c.mu, c.sigma, c.tolerance, c.want, got)
		}
	}
}
