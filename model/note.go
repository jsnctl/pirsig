package model

type Note struct {
	Seed     float64 `yaml:"seed"`
	Duration float64 `yaml:"duration"`
	WaveFn   string  `yaml:"wave"`
	Decay    float64 `yaml:"decay"`
}
