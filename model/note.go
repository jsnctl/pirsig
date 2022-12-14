package model

import "github.com/jsnctl/pirsig/waveforms"

type Note struct {
	Seed      interface{} `yaml:"seed"`
	Duration  interface{} `yaml:"duration"`
	Waveform  interface{} `yaml:"waveform"`
	Decay     float64     `yaml:"decay"`
	Reverb    bool        `yaml:"reverb"`
	Amplitude float64     `yaml:"amplitude"`
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

func (n *Note) GetWaveform() func(float64, float64) float64 {
	waveform := waveforms.TryWaveform(n.Waveform)
	if waveform != nil {
		return waveforms.WaveformLookup[waveform.Type]
	}

	switch n.Waveform.(type) {
	case string:
		return waveforms.WaveformLookup[n.Waveform.(string)]
	}

	return nil
}
