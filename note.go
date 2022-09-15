package main

import "github.com/jsnctl/pirsig/shared"

type Note struct {
	wave []float64
}

func createNote(seed float64, duration float64, waveFn func(float64, float64) float64) Note {
	note := Note{}
	nSamples := int(duration * shared.SampleRate)

	angleIncrement := shared.Tau / float64(nSamples)

	for i := 0; i <= nSamples; i++ {
		angle := angleIncrement * float64(i)
		sample := waveFn(angle, seed)
		note.wave = append(note.wave, sample)
	}
	return note
}
