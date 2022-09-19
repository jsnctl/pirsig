package test

import (
	"github.com/jsnctl/pirsig/io"
	"github.com/jsnctl/pirsig/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWritePolyphonic(t *testing.T) {
	left := model.Wave{}
	right := model.Wave{}

	left.Values = []float64{1, 1, 1, 1}
	right.Values = []float64{0, 0, 0, 0}

	expected := []float64{1, 0, 1, 0}

	assert.Equal(t, expected, io.Polyphonic(left, right).Values)
}
