package response

// RecordingItem represent recording data model
type RecordingItem struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	Emotion   string `json:"emotion"`
	IsActive  bool   `json:"is_active"`
	AudioSrc  string `json:"audio_src,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Recording represent recording response data model
type Recording struct {
	Success bool          `json:"success"`
	Item    RecordingItem `json:"item"`
}

// Recordings represent recordings response data model
type Recordings struct {
	Success  bool            `json:"success"`
	Page     int             `json:"page"`
	NumPages int             `json:"num_pages"`
	PageSize int             `json:"page_size"`
	Items    []RecordingItem `json:"items"`
}
