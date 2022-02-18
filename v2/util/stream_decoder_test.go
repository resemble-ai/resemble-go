package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlushBufferWithHeader(t *testing.T) {
	decoder, err := NewStreamDecoder(2, false)
	assert.NoError(t, err)
	data := []byte("test data byte test data")
	decoder.DecodeChunk(data)

	result := []byte{}
	for {
		if b := decoder.FlushBuffer(false); b != nil {
			assert.LessOrEqual(t, len(b), 2)
			result = append(result, b...)
		} else {
			break
		}
	}

	assert.Equal(t, len(data), len(result))
}

func TestFlushBufferWithoutHeader(t *testing.T) {
	decoder, err := NewStreamDecoder(4, true)
	assert.NoError(t, err)
	data := []byte("contain data contain data contain data contain data contain data data is ready.")
	decoder.DecodeChunk(data)

	result := []byte{}
	for {
		if b := decoder.FlushBuffer(false); b != nil {
			assert.LessOrEqual(t, len(b), 4)
			result = append(result, b...)
		} else {
			break
		}
	}

	assert.Equal(t, (len(data) - streamingWavHeaderBufferLength), len(result))
}

func TestFlushBufferWithForce(t *testing.T) {
	decoder, err := NewStreamDecoder(2, false)
	assert.NoError(t, err)
	data := []byte("test data byte test data")
	decoder.DecodeChunk(data)

	b := decoder.FlushBuffer(true)
	assert.Equal(t, data, b)
}
