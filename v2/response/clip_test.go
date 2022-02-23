package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetaData(t *testing.T) {
	data := []byte{82, 73, 70, 70, 228, 127, 1, 0, 87, 65, 86, 69, 102, 109, 116, 32, 16, 0, 0, 0, 1, 0, 1, 0, 0, 125, 0, 0, 0, 250, 0, 0, 2, 0, 16, 0, 99, 117, 101, 32, 20, 2, 0, 0}
	meta := NewMetaData(data)
	assert.Equal(t, data, meta.RawData)
	assert.Equal(t, "RIFF", meta.RiffID)
	assert.Greater(t, meta.FileSize, 0)
	assert.Equal(t, "WAVE", meta.RiffType)
	assert.Equal(t, "fmt ", meta.FormatChunkID)
	assert.Equal(t, 16, meta.ChunkDataSize)
	assert.Equal(t, 1, meta.CompressionCode)
	assert.Equal(t, 1, meta.NumberOfChannels)
	assert.True(t, (meta.SampleRate >= 8000 && meta.SampleRate <= 48000))
	assert.True(t, (meta.ByteRate >= 16000 && meta.ByteRate <= 96000))

	assert.Equal(t, 2, meta.BlockAlign)
	assert.Equal(t, 16, meta.BitsPerSample)
}
