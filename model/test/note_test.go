package test

import (
	"github.com/jsnctl/pirsig/model"
	"github.com/jsnctl/pirsig/shared"
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
