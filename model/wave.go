package model

import (
	"github.com/jsnctl/pirsig/shared"
	"github.com/jsnctl/pirsig/waveforms"
	"math"
)

type Wave struct {
	Values []float64
}

const Delay = 2000
const Decay = 0.8

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

	if note.Reverb { // needs review
		var ghost []float64
		ghost = wave.Values
		for i := range wave.Values {
			if i == 0 || i >= len(wave.Values)-1 || i < Delay {
				continue
			}
			wave.Values[i] = 0.5 * (wave.Values[i] + (Decay * ghost[i-(Delay-1)]))
		}
	}

	return wave
}
