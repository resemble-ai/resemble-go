package response

import (
	"bytes"
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

// MetaHeader represent Header & Format chunks
type MetaHeader struct {
	// RiffID  Riff ID
	RiffID string
	// RiffType Riff type
	RiffType string
	// FileSize File Size
	FileSize int
	// FormatChunkID Format chunk id
	FormatChunkID string
	// ChunkDataSize Chunk data size
	ChunkDataSize int
	// CompressionCode Compression code
	CompressionCode int
	// NumberOfChannels number of channel
	NumberOfChannels int
	// SampleRate Sample rate
	SampleRate int
	// ByteRate Byte rate
	ByteRate int
	// BlockAlign  Block align
	BlockAlign int
	// BitsPerSample  Bits per sample
	BitsPerSample int
}

type Cue struct {
	// CueChunkID Cue chunk ID
	CueChunkID string
	// RemSizeCue Remaining size of the cue chunk after this read
	RemSizeCue int
	// NumberCuePoint Number of remaining cue points
	NumberCuePoint int
}

type CuePoint struct {
	// CuePointsID 	Cue point ID
	CuePointsID uint32
	// SamplesOffset Sample offset
	SamplesOffset uint32
}

type List struct {
	// ListChunkID 	List chunk ID
	ListChunkID string
	// RemSizeofListChunk Remaining size of the list chunk after this read
	RemSizeofListChunk int
	// TypeID Type ID
	TypeID string
}

type Ltxt struct {
	// LtxtChunkID 	LTXT chunk ID
	LtxtChunkID string

	// RemsizeLtxtChunk Remaining size of this ltxt chunk after this read*
	RemsizeLtxtChunk int

	// LtxtCuePointID  Cue point ID
	LtxtCuePointID uint32

	// LengthNumSample Length in number of samples
	LengthNumSample uint32

	// CharType Character type
	CharType string

	// TextLength
	TextLength string
}

// MetaTimestamps
type MetaTimestamps struct {
	// Cue Cue
	Cue Cue
	// CuePoints Cue points
	CuePoints []CuePoint
	// List list chunk
	List List
	// Ltxt ltxt chunks
	Ltxt []Ltxt
}

type AudioDataChunk struct {
	// DataChunkID Data chunk ID
	DataChunkID string
	// NumberOfRemAudioSamples Number of remaining audio samples * 2
	NumberOfRemAudioSamples int
}

// Metadata represent metadata of wav file
type Metadata struct {
	// Header represent Header & Format chunks
	Header MetaHeader
	// Timestamps (cue, list & ltxt chunks)
	TimeStamps MetaTimestamps

	// AudioData Audio data chunk
	AudioData AudioDataChunk

	// RawData contains raw metadata
	rawData []byte

	foundLtxt bool
	foundData bool
	ltxtIndex int
	dataIndex int
	isFlush   bool
}

func NewMetaData() *Metadata {
	return &Metadata{}
}

func (meta *Metadata) setFoundLtxt() bool {
	index := bytes.LastIndex(meta.rawData, []byte("ltxt"))
	if index > -1 {
		meta.foundLtxt = true
		meta.ltxtIndex = index
	}
	return meta.foundLtxt
}

func (meta *Metadata) setFoundData() bool {
	if !meta.setFoundLtxt() {
		return false
	}
	index := bytes.LastIndex(meta.rawData, []byte("data"))
	if index > meta.ltxtIndex {
		meta.foundData = true
		meta.dataIndex = index
	}
	return meta.foundData
}

// GetRawData returns raw metadata
func (m Metadata) GetRawData() []byte {
	return m.rawData
}

func (meta *Metadata) Flush(data []byte) *Metadata {
	if meta.isFlush {
		return nil
	}

	meta.rawData = append(meta.rawData, data...)

	if !meta.setFoundData() {
		return nil
	}

	if m := meta.generate(); m != nil {
		meta.isFlush = true
		return m
	}

	return nil
}

// generate returns new instance of metadata
func (meta *Metadata) generate() *Metadata {
	if len(meta.rawData) < (meta.dataIndex + 8) {
		return nil
	}
	data := meta.rawData[0:(meta.dataIndex + 8)]
	meta.rawData = data

	meta.Header.RiffID = string(data[0:4])
	meta.Header.FileSize = (bitsToInt(data[4:8]))
	meta.Header.RiffType = string(data[8:12])
	meta.Header.FormatChunkID = string(data[12:16])
	meta.Header.ChunkDataSize = bitsToInt(data[16:20])
	meta.Header.CompressionCode = bitsToInt(data[20:22])
	meta.Header.NumberOfChannels = bitsToInt(data[22:24])
	meta.Header.SampleRate = bitsToInt(data[24:28])
	meta.Header.ByteRate = bitsToInt(data[28:32])
	meta.Header.BlockAlign = bitsToInt(data[32:34])
	meta.Header.BitsPerSample = bitsToInt(data[34:36])

	// Timestamps (cue, list & ltxt chunks)#
	meta.TimeStamps.Cue.CueChunkID = string(data[36:40])
	meta.TimeStamps.Cue.RemSizeCue = bitsToInt(data[40:44])
	meta.TimeStamps.Cue.NumberCuePoint = bitsToInt(data[44:48])

	position := 48
	for i := 0; i < meta.TimeStamps.Cue.NumberCuePoint; i++ {
		meta.TimeStamps.CuePoints = append(meta.TimeStamps.CuePoints, CuePoint{
			CuePointsID:   uint32(bitsToInt64(data[position : position+4])),
			SamplesOffset: uint32(bitsToInt64(data[position+20 : position+24])),
		})
		position = position + 24
	}

	meta.TimeStamps.List.ListChunkID = string(data[position : position+4])
	position = position + 4
	meta.TimeStamps.List.RemSizeofListChunk = bitsToInt(data[position : position+4])
	position = position + 4
	meta.TimeStamps.List.TypeID = string(data[position : position+4])
	position = position + 4

	for {
		testLtxt := string(data[position : position+4])
		if testLtxt != "ltxt" {
			break
		}
		ltxt := Ltxt{}
		ltxt.LtxtChunkID = testLtxt
		position = position + 4
		ltxt.RemsizeLtxtChunk = bitsToInt(data[position : position+4])
		position = position + 4
		ltxt.LtxtCuePointID = uint32(bitsToInt64(data[position : position+4]))
		position = position + 4
		ltxt.LengthNumSample = uint32(bitsToInt64(data[position : position+4]))
		position = position + 4
		ltxt.CharType = string(data[position : position+4])
		position = position + 4 + 8
		textLength := ltxt.RemsizeLtxtChunk - 20
		skipLength := textLength
		if skipLength%2 != 0 {
			skipLength = skipLength - 1
			textLength = textLength + 1
		}
		b := data[position : position+skipLength]
		ltxt.TextLength = string(b)
		meta.TimeStamps.Ltxt = append(meta.TimeStamps.Ltxt, ltxt)
		position = position + textLength
	}

	meta.AudioData.DataChunkID = string(data[position : position+4])
	position = position + 4
	meta.AudioData.NumberOfRemAudioSamples = bitsToInt(data[position : position+4])

	return meta
}

func bitsToInt(b []byte) int {
	return int(bitsToInt64(b))
}

func bitsToInt64(b []byte) uint64 {
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
	return bits
}
