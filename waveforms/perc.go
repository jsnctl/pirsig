package waveforms

import (
	"math"
	"math/rand"
)

func Kick(angle float64, frequency float64) float64 {
	adjusted := 10 * math.Exp(-angle/frequency)
	return 5 * Sine(angle, adjusted)
}

func Snare(angle float64, frequency float64) float64 {
	sweepFrequency := 10 * frequency * math.Exp(-angle)
	return 0.5 * Sine(angle, sweepFrequency) * float64(rand.Intn(50)) * math.Exp(-angle/3)
}
