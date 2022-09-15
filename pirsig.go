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

	for _, note := range s.Notes {
		wave := model.CreateWave(note)
		io.Monophonic(wave)
	}
}
