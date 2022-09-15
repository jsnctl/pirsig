package main

import (
	"encoding/binary"
	"math"
	"os"
)

func monophonic(note Note) {
	for _, value := range note.wave {
		writeNoteValue(value, f)
	}
}

func writeNoteValue(value float64, file *os.File) {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(value)))
	_, _ = file.Write(buf[:])
}
