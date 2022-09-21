package test

import (
	"github.com/jsnctl/pirsig/waveforms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaveformsAdditive(t *testing.T) {
	Init()
	fns := []func(float64, float64) float64{
		waveforms.Sine,
		waveforms.Triangle,
	}

	result := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		angle := float64(i) * angleIncrement
		result[i] = waveforms.Additive(fns, angle, frequency)
	}

	assert.NotNil(t, result)
}
