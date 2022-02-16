package option

const (
	streamingBufferSize = 4 * 1024
	streamingChunkSize  = 2
)

// ClipStream to hold clip stream option data model
type ClipStream struct {
	BufferSize    int
	ChunkSize     int
	WithWavHeader bool
}

// Parse format option with proper value
func (c *ClipStream) Parse() {
	if c.BufferSize < 1 {
		c.BufferSize = streamingBufferSize
	}
	if c.ChunkSize < 1 {
		c.ChunkSize = streamingChunkSize
	}
}
