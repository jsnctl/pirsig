package main

import (
	"github.com/jsnctl/pirsig/io"
	"github.com/jsnctl/pirsig/model"
	"github.com/jsnctl/pirsig/shared"
	"os"
)

func main() {
	io.F, _ = os.Create(shared.OutputFile)
	s := io.Read()

	for _, sequence := range s.Sequences {
		if len(sequence.Tracks) > 1 {
			left := model.Arp(sequence.Length, sequence.Tracks[0].Notes)
			right := model.Arp(sequence.Length, sequence.Tracks[1].Notes)
			for i := range left {
				leftWave := model.CreateWave(left[i])
				rightWave := model.CreateWave(right[i])
				combined := io.Polyphonic(leftWave, rightWave)
				io.Monophonic(model.LowPass(combined, 1000))
				//io.Monophonic(combined)
			}
		} else {
			notes := model.Arp(sequence.Length, sequence.Tracks[0].Notes)
			for _, note := range notes {
				wave := model.CreateWave(note)
				io.Monophonic(wave)
			}
		}
	}

}
