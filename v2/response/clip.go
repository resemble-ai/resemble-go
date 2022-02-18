package response

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

// ClipStream represent clip stream response data model
type ClipStream struct {
	Chunk []byte
	Err   error
}
