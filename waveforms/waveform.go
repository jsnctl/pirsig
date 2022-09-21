package waveforms

import "github.com/mitchellh/mapstructure"

type Waveform struct {
	Type string `yaml:"type"`
}

func TryWaveform(i interface{}) *Waveform {
	waveform := Waveform{}
	err := mapstructure.Decode(i, &waveform)
	if err != nil {
		return nil
	}
	return &waveform
}

var WaveformLookup = map[string]func(float64, float64) float64{
	"additive": DummyAdditive,
	"sine":     Sine,
	"square":   Square,
	"triangle": Triangle,
	"saw":      Sawtooth,
	"kick":     Kick,
	"snare":    Snare,
}
