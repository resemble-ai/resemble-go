package response

import (
	"encoding/binary"
	"log"
)

type TimeStamps struct {
	GraphChars []string    `json:"graph_chars,omitempty"`
	GraphTimes [][]float64 `json:"graph_times,omitempty"`
	PhonChars  []string    `json:"phon_chars,omitempty"`
	PhonTimes  [][]float64 `json:"phon_times,omitempty"`
}

// ClipItem represent clip data model
type ClipItem struct {
	UUID       string      `json:"uuid"`
	Title      string      `json:"title"`
	Body       string      `json:"body"`
	VoiceUUID  string      `json:"voice_uuid"`
	IsPublic   bool        `json:"is_public"`
	IsArchived bool        `json:"is_archived"`
	TimeStamps TimeStamps  `json:"timestamps,omitempty"`
	AudioSrc   string      `json:"audio_src"`
	RawAudio   interface{} `json:"raw_audio,omitempty"`
	CreatedAt  string      `json:"created_at"`
	UpdatedAt  string      `json:"updated_at"`
}

// Clip represent clip response data model
type Clip struct {
	Success bool     `json:"success"`
	Item    ClipItem `json:"item"`
}

// Clips represent clips response data model
type Clips struct {
	Success  bool       `json:"success"`
	Page     int        `json:"page"`
	NumPages int        `json:"num_pages"`
	PageSize int        `json:"page_size"`
	Items    []ClipItem `json:"items"`
}

// Metadata represent metadata of wav file
type Metadata struct {
	// Riff ID
	RiffID string
	// Riff type
	RiffType string
	// File Size
	FileSize int
	// Format chunk id
	FormatChunkID string
	// Chunk data size
	ChunkDataSize int
	// Compression code
	CompressionCode int
	// number of channel
	NumberOfChannels int
	// Sample rate
	SampleRate int
	// Byte rate
	ByteRate int
	// Block align
	BlockAlign int
	// Bits per sample
	BitsPerSample int
	// RawData contains raw metadata
	RawData []byte
}

// NewMetaData returns new instance of metadata
func NewMetaData(data []byte) Metadata {
	meta := Metadata{
		RawData: data,
	}
	if len(data) != 44 {
		log.Println("data lenght must be 44")
		return meta
	}

	meta.RiffID = string(data[0:4])
	meta.FileSize = (bitsToInt(data[4:8]) - 8)
	meta.RiffType = string(data[8:12])
	meta.FormatChunkID = string(data[12:16])
	meta.ChunkDataSize = bitsToInt(data[16:20])
	meta.CompressionCode = bitsToInt(data[20:22])
	meta.NumberOfChannels = bitsToInt(data[22:24])
	meta.SampleRate = bitsToInt(data[24:28])
	meta.ByteRate = bitsToInt(data[28:32])
	meta.BlockAlign = bitsToInt(data[32:34])
	meta.BitsPerSample = bitsToInt(data[34:36])

	return meta
}

func bitsToInt(b []byte) int {
	var bits uint64
	switch len(b) {
	case 2:
		bits = uint64(binary.LittleEndian.Uint16(b))
	case 4:
		bits = uint64(binary.LittleEndian.Uint32(b))
	case 8:
		bits = binary.LittleEndian.Uint64(b)
	default:
		log.Println("Can't parse to int ")
	}
	return int(bits)
}
