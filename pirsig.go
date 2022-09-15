package main

import (
	"github.com/jsnctl/pirsig/shared"
	"github.com/jsnctl/pirsig/waveforms"
	"os"
)

var f *os.File

func main() {
	f, _ = os.Create(shared.OutputFile)

	values := []float64{400, 300, 200, 400, 300, 200}

	for _, value := range values {
		note := createNote(value, 0.2, waveforms.Sine)
		monophonic(note)
	}
}
