package model

import (
	"github.com/jsnctl/pirsig/shared"
	"github.com/jsnctl/pirsig/waveforms"
	"math"
)

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
		decayFactor := 1.0
		if note.Decay != 0 {
			decayFactor = math.Exp(-angle / note.Decay)
		}
		sample := waveFunction(angle, note.Seed) * decayFactor
		wave.Values = append(wave.Values, sample)
	}
	return wave
}
