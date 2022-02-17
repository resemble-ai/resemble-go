package response

// Message represent generic message response data model
type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
