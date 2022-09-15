package io

import (
	"encoding/binary"
	"github.com/jsnctl/pirsig/model"
	"math"
	"os"
)

var F *os.File

func Monophonic(wave model.Wave) {
	for _, value := range wave.Values {
		writeNoteValue(value, F)
	}
}

func writeNoteValue(value float64, file *os.File) {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(value)))
	_, _ = file.Write(buf[:])
}
