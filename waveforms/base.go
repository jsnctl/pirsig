package waveforms

import "math"

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}
