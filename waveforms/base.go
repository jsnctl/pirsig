package waveforms

import (
	"github.com/jsnctl/pirsig/shared"
	"math"
)

var Lookup = map[string]func(float64, float64) float64{
	"sine":     Sine,
	"square":   Square,
	"triangle": Triangle,
	"saw":      Sawtooth,
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}

func Triangle(angle float64, frequency float64) float64 {
	sineValue := Sine(angle, frequency)
	return (4 / (shared.Tau)) * math.Asin(sineValue)
}

func Sawtooth(angle float64, frequency float64) float64 {
	cotValue := math.Tan(angle * frequency)
	return (2 / math.Pi) * math.Atan(1.0/cotValue)
}
