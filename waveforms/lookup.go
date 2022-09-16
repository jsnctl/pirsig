package waveforms

var WaveformLookup = map[string]func(float64, float64) float64{
	"sine":     Sine,
	"square":   Square,
	"triangle": Triangle,
	"saw":      Sawtooth,
	"kick":     Kick,
	"snare":    Snare,
}
