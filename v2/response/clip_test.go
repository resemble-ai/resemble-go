package response

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMetaData(t *testing.T) {
	r := loadTestData()
	defer r.Close()

	streamingBufferSize := 4 * 1024
	buf := make([]byte, streamingBufferSize)
	meta := NewMetaData()
	rawData := make([]byte, 0)
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		b := buf[:n]
		rawData = append(rawData, b...)
		if m := meta.Flush(b); m != nil {
			meta = m
		}
	}

	index := bytes.LastIndex(rawData, []byte("data"))
	assert.Equal(t, (index + 8), len(meta.GetRawData()))
	assert.Equal(t, "RIFF", meta.Header.RiffID)
	assert.Greater(t, meta.Header.FileSize, 0)
	assert.Equal(t, "WAVE", meta.Header.RiffType)
	assert.Equal(t, "fmt ", meta.Header.FormatChunkID)
	assert.Equal(t, 16, meta.Header.ChunkDataSize)
	assert.Equal(t, 1, meta.Header.CompressionCode)
	assert.Equal(t, 1, meta.Header.NumberOfChannels)
	assert.True(t, (meta.Header.SampleRate >= 8000 && meta.Header.SampleRate <= 48000))
	assert.True(t, (meta.Header.ByteRate >= 16000 && meta.Header.ByteRate <= 96000))

	assert.Equal(t, 2, meta.Header.BlockAlign)
	assert.Equal(t, 16, meta.Header.BitsPerSample)

	assert.Equal(t, "cue ", meta.TimeStamps.Cue.CueChunkID)
	assert.Equal(t, 4+(meta.TimeStamps.Cue.NumberCuePoint*24), meta.TimeStamps.Cue.RemSizeCue)
	assert.Greater(t, meta.TimeStamps.Cue.NumberCuePoint, 1)

	assert.Equal(t, "list", meta.TimeStamps.List.ListChunkID)
	assert.Greater(t, meta.TimeStamps.List.RemSizeofListChunk, 1)
	assert.Equal(t, "adtl", meta.TimeStamps.List.TypeID)

	for _, ltxt := range meta.TimeStamps.Ltxt {
		assert.Equal(t, "ltxt", ltxt.LtxtChunkID)
		assert.True(t, ltxt.CharType == "grph" || ltxt.CharType == "phon")
	}

	assert.Equal(t, "data", meta.AudioData.DataChunkID)
}
