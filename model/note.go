package model

type Note struct {
	Seed     float64       `yaml:"seed"`
	Dist     *Distribution `yaml:"dist"`
	Duration float64       `yaml:"duration"`
	WaveFn   string        `yaml:"wave"`
	Decay    float64       `yaml:"decay"`
	Reverb   bool          `yaml:"reverb"`
}
