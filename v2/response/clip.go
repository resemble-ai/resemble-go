package response

// ClipItem represent clip data model
type ClipItem struct {
	UUID       string `json:"uuid"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	VoiceUUID  string `json:"voice_uuid"`
	IsPublic   bool   `json:"is_public"`
	IsArchived bool   `json:"is_archived"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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
