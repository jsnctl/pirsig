package waveforms

func Additive(
	fns []func(float64, float64) float64,
	angle float64, frequency float64,
) float64 {
	var value float64
	for _, fn := range fns {
		value = value + fn(angle, frequency)
	}
	return value
}

func DummyAdditive(angle float64, frequency float64) float64 {
	fns := []func(float64, float64) float64{
		Sawtooth,
		Kick,
	}

	return Additive(fns, angle, frequency)
}
