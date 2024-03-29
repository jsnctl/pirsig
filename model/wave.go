package model

import (
	"github.com/jsnctl/pirsig/shared"
	"math"
)

type Wave struct {
	Values []float64
}

const Delay = 2000
const Decay = 0.8

func CreateWave(note Note) Wave {
	wave := Wave{}
	seed := note.GetSeed()
	duration := note.GetDuration()
	waveform := note.GetWaveform()

	if note.Amplitude == 0 {
		note.Amplitude = 1
	}

	nSamples := int(duration * shared.SampleRate)
	angleIncrement := shared.Tau / float64(nSamples)

	for i := 0; i <= nSamples; i++ {
		angle := angleIncrement * float64(i)
		decayFactor := 1.0
		if note.Decay != 0 {
			decayFactor = math.Exp(-angle / note.Decay)
		}

		sample := note.Amplitude * waveform(angle, seed) * decayFactor
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

func LowPass(wave Wave, freq float64) Wave {
	K := math.Tan((math.Pi * freq) / shared.SampleRate)
	b0 := K / (K + 1)
	b1 := K / (K + 1)
	a1 := (K - 1) / (K + 1)
	for i := range wave.Values {
		if i == 0 || i >= len(wave.Values) {
			continue
		}
		wave.Values[i] = wave.Values[i] - a1*wave.Values[i-1]
		wave.Values[i] = (b0 * wave.Values[i]) + (b1 * wave.Values[i-1])
	}
	return wave
}
