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

func Polyphonic(left model.Wave, right model.Wave) model.Wave {
	combined := model.Wave{}
	for i := 0; i < min(len(left.Values), len(right.Values)); i += 2 {
		combined.Values = append(combined.Values, left.Values[i])
		combined.Values = append(combined.Values, right.Values[i])
	}
	return combined
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func writeNoteValue(value float64, file *os.File) {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(value)))
	_, _ = file.Write(buf[:])
}
