package model

type Sequence struct {
	Length int     `yaml:"length"`
	Tracks []Track `yaml:"tracks"`
}

type Track struct {
	Notes []Note `yaml:"notes"`
}

func Arp(length int, pattern []Note) []Note {
	sequence := make([]Note, length)
	cursor := 0
	for i := range sequence {
		sequence[i] = pattern[cursor]
		if cursor < len(pattern)-1 {
			cursor += 1
		} else {
			cursor = 0
		}
	}
	return sequence
}
