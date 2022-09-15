package model

import (
	"github.com/jsnctl/pirsig/shared"
	"github.com/jsnctl/pirsig/waveforms"
)

type Note struct {
	Seed     float64 `yaml:"Seed"`
	Duration float64 `yaml:"Duration"`
	WaveFn   string  `yaml:"WaveFn"`
}

type Wave struct {
	Values []float64
}

func CreateWave(note Note) Wave {
	wave := Wave{}
	nSamples := int(note.Duration * shared.SampleRate)
	waveFunction := waveforms.Lookup[note.WaveFn]

	angleIncrement := shared.Tau / float64(nSamples)

	for i := 0; i <= nSamples; i++ {
		angle := angleIncrement * float64(i)
		sample := waveFunction(angle, note.Seed)
		wave.Values = append(wave.Values, sample)
	}
	return wave
}
