package model

type Note struct {
	Seed     interface{} `yaml:"seed"`
	Duration interface{} `yaml:"duration"`
	WaveFn   string      `yaml:"wave"`
	Decay    float64     `yaml:"decay"`
	Reverb   bool        `yaml:"reverb"`
}

func (n *Note) GetSeed() float64 {
	dist := TryDistribution(n.Seed)
	if dist != nil {
		opts := DistributionOpts{mu: dist.Mu, sigma: dist.Sigma, choices: dist.Choices}
		return DistributionLookup[dist.Type](opts)
	}

	switch n.Seed.(type) {
	case int:
		return float64(n.Seed.(int))
	case float64:
		return n.Seed.(float64)
	}

	return 0
}

func (n *Note) GetDuration() float64 {
	dist := TryDistribution(n.Duration)
	if dist != nil {
		opts := DistributionOpts{mu: dist.Mu, sigma: dist.Sigma, choices: dist.Choices}
		return DistributionLookup[dist.Type](opts)
	}

	switch n.Duration.(type) {
	case int:
		return float64(n.Duration.(int))
	case float64:
		return n.Duration.(float64)
	}

	return 0
}
