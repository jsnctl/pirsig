package waveforms

import (
	"github.com/jsnctl/pirsig/shared"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var nSamples int
var frequency float64
var angleIncrement float64

const testSampleRate = 44000

func Init() {
	nSamples = testSampleRate * 1
	frequency = 1000.0
	angleIncrement = (2 * math.Pi) / float64(nSamples)
}

func waveformHarness(waveform func(float64, float64) float64) []float64 {
	result := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		angle := float64(i) * angleIncrement
		result[i] = waveform(angle, frequency)
	}
	return result
}

func TestSine(t *testing.T) {
	Init()
	result := waveformHarness(Sine)

	expectedFirstTen := []float64{
		0, 0.14231483827328514, 0.28173255684142967, 0.4154150130018864,
		0.5406408174555976, 0.6548607339452851, 0.7557495743542582,
		0.8412535328311811, 0.9096319953545183, 0.9594929736144975,
	}

	for i, expected := range expectedFirstTen {
		assert.True(t,
			shared.FloatingPointEqual(
				expected,
				result[i],
				1e-10,
			),
		)
	}
}

func TestSquare(t *testing.T) {
	Init()
	result := waveformHarness(Square)

	assert.Equal(t, 1.0, result[0])
	assert.Equal(t, -1.0, result[23])
	assert.Equal(t, 1.0, result[45])
	assert.Equal(t, -1.0, result[67])
}

func TestTriangle(t *testing.T) {
	Init()
	result := waveformHarness(Triangle)

	// not perfect assertions for triangle wave
	expectedFirstTen := []float64{
		0, 0.0909090909090909, 0.18181818181818182, 0.27272727272727276,
		0.36363636363636365, 0.4545454545454546, 0.5454545454545454,
		0.6363636363636362, 0.7272727272727272, 0.8181818181818183,
	}

	for i, expected := range expectedFirstTen {
		assert.True(t,
			shared.FloatingPointEqual(
				expected,
				result[i],
				1e-10,
			),
		)
	}

	assert.Equal(t, 1.0, result[11])
}
