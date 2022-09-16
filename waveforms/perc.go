package waveforms

import "math"

func Kick(angle float64, frequency float64) float64 {
	adjusted := 10 * math.Exp(-angle/frequency)
	return 5 * Sine(angle, adjusted)
}
