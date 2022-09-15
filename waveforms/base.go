package waveforms

import "math"

var Lookup = map[string]func(float64, float64) float64{
	"sine": Sine,
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}
