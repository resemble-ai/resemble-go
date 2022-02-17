package response

// ProjectItem represent project data model
type ProjectItem struct {
	UUID            string `json:"uuid"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	IsPublic        bool   `json:"is_public"`
	IsCollaborative bool   `json:"is_collaborative"`
	IsArchived      bool   `json:"is_archived"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

// Project represent project response data model
type Project struct {
	Success bool        `json:"success"`
	Item    ProjectItem `json:"item"`
}

// Projects represent projects response data model
type Projects struct {
	Success  bool          `json:"success"`
	Page     int           `json:"page"`
	NumPages int           `json:"num_pages"`
	PageSize int           `json:"page_size"`
	Items    []ProjectItem `json:"items"`
}
