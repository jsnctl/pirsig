package model

import "math/rand"

type Distribution struct {
	Type  string  `yaml:"type"`
	Mu    float64 `yaml:"mu"`
	Sigma float64 `yaml:"sigma"`
}

var DistributionLookup = map[string]func(float64, float64) float64{
	"gaussian": Gaussian,
}

func Gaussian(mu float64, sigma float64) float64 {
	return rand.NormFloat64()*sigma + mu
}
