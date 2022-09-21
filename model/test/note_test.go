package test

import (
	"github.com/jsnctl/pirsig/model"
	"github.com/jsnctl/pirsig/shared"
	"github.com/jsnctl/pirsig/waveforms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoteGetSeed(t *testing.T) {
	note := model.Note{
		Seed: 10,
	}
	assert.Equal(t, 10.0, note.GetSeed())

	note = model.Note{
		Seed: model.Distribution{
			Type:  "gaussian",
			Mu:    10,
			Sigma: 0.5,
		},
	}
	assert.True(t, shared.FloatingPointEqual(10, note.GetSeed(), 1))
}

func TestNoteGetWaveform(t *testing.T) {
	note := model.Note{
		Waveform: "saw",
	}
	// https://github.com/stretchr/testify/issues/565
	// comparison of func not supported
	assert.NotNil(t, note.GetWaveform())

	note = model.Note{
		Waveform: waveforms.Waveform{
			Type: "saw",
		},
	}
	assert.NotNil(t, note.GetWaveform())
}
