package util

import "fmt"

const streamingWavHeaderBufferLength = 44

// StreamDecoder represent StreamDecoder model
type StreamDecoder struct {
	bufferSize      int
	ignoreWavHeader bool
	headerBuffer    []byte
	buffer          []byte
}

// NewStreamDecoder returns a new instance of StreamDecoder
func NewStreamDecoder(bufferSize int, ignoreWavHeader bool) (*StreamDecoder, error) {
	if bufferSize < 2 {
		return nil, fmt.Errorf("buffer size cannot be less than 2")
	}
	if bufferSize%2 != 0 {
		return nil, fmt.Errorf("buffer size must be evenly divisible by 2")
	}

	return &StreamDecoder{
		bufferSize:      bufferSize,
		ignoreWavHeader: ignoreWavHeader,
	}, nil
}

// DecodeChunk decode chunk data
func (s *StreamDecoder) DecodeChunk(chunk []byte) {
	if (len(s.headerBuffer) < streamingWavHeaderBufferLength) && s.ignoreWavHeader {
		s.headerBuffer = append(s.headerBuffer, chunk...)
		if len(s.headerBuffer) >= streamingWavHeaderBufferLength {
			s.buffer = s.headerBuffer[streamingWavHeaderBufferLength:]
			s.headerBuffer = s.headerBuffer[0:streamingWavHeaderBufferLength]
		}
	} else {
		s.buffer = append(s.buffer, chunk...)
	}
}

// FlushBuffer returns chunk data
func (s *StreamDecoder) FlushBuffer(force bool, drain ...bool) []byte {
	if force {
		defer s.Reset()
		return s.buffer
	}

	if len(s.buffer) >= s.bufferSize {
		buff := s.buffer[0:s.bufferSize]
		s.buffer = s.buffer[s.bufferSize:]
		return buff
	} else {
		if len(drain) > 0 && drain[0] {
			defer s.Reset()
			return s.buffer
		}
	}

	return nil
}

// Reset reset the buffer variable
func (s *StreamDecoder) Reset() {
	s.buffer = nil
}

// Header returns wav header value
func (s *StreamDecoder) Header() []byte {
	if s.ignoreWavHeader {
		return s.headerBuffer
	}

	return s.buffer[:streamingWavHeaderBufferLength]
}
