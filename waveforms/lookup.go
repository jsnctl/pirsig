package waveforms

var Lookup = map[string]func(float64, float64) float64{
	"sine":     Sine,
	"square":   Square,
	"triangle": Triangle,
	"saw":      Sawtooth,
	"kick":     Kick,
	"snare":    Snare,
}
