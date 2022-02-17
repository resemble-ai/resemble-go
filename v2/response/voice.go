package response

// VoiceItem represent voice data model
type VoiceItem struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	IsPublic        bool   `json:"is_public"`
	IsCollaborative bool   `json:"is_collaborative"`
	IsArchived      bool   `json:"is_archived"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// Voice represent voice response data model
type Voice struct {
	Success bool      `json:"success"`
	Item    VoiceItem `json:"item"`
}

// Voices represent voices response data model
type Voices struct {
	Success  bool        `json:"success"`
	Page     int         `json:"page"`
	NumPages int         `json:"num_pages"`
	PageSize int         `json:"page_size"`
	Items    []VoiceItem `json:"items"`
}
