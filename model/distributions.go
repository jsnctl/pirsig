package model

import (
	"github.com/mitchellh/mapstructure"
	"math/rand"
)

type Distribution struct {
	Type    string    `yaml:"type"`
	Mu      float64   `yaml:"mu"`
	Sigma   float64   `yaml:"sigma"`
	Choices []float64 `yaml:"choices"`
}

func TryDistribution(i interface{}) *Distribution {
	dist := Distribution{}
	err := mapstructure.Decode(i, &dist)
	if err != nil {
		return nil
	}
	return &dist
}

var DistributionLookup = map[string]func(opts DistributionOpts) float64{
	"gaussian": Gaussian,
	"choice":   Choice,
}

type DistributionOpts struct {
	mu      float64
	sigma   float64
	choices []float64
}

func Gaussian(opts DistributionOpts) float64 {
	return rand.NormFloat64()*opts.sigma + opts.mu
}

func Choice(opts DistributionOpts) float64 {
	return opts.choices[rand.Intn(len(opts.choices))]
}
